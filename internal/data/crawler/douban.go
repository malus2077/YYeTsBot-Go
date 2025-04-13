package crawler

import (
	"YYeTsBot-Go/internal/biz"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/gocolly/colly/v2"
)

const (
	DoubanSearch = "https://www.douban.com/search?cat=1002&q=%s"
	DoubanDetail = "https://movie.douban.com/subject/%d"
)

type Douban struct {
	logger *log.Helper
}

func NewDouban() *Douban {
	return &Douban{logger: log.NewHelper(log.DefaultLogger)}
}

func (d *Douban) FetchMediaInfo(name string) *biz.Douban {
	doubanID, err := d.extraIdByName(name)
	if err != nil {
		d.logger.Errorf("提取豆瓣ID失败: %v", err)
		return nil
	}
	mediaInfo := d.extraInfoById(doubanID)
	return mediaInfo
}

func (d *Douban) extraIdByName(name string) (int64, error) {
	c := colly.NewCollector()
	searchURL := fmt.Sprintf(DoubanSearch, url.QueryEscape(name))
	var doubanID int64
	var found bool

	c.OnHTML("div.content", func(e *colly.HTMLElement) {
		if found { // 如果已经找到，则停止后续处理
			return
		}
		link := e.ChildAttr("a", "href")
		if link == "" {
			return
		}

		decodedLink, err := url.QueryUnescape(link)
		if err != nil {
			d.logger.Errorf("URL解码失败: %v", err)
			return
		}

		re := regexp.MustCompile(`https://movie\.douban\.com/subject/(\d+)/`)
		matches := re.FindStringSubmatch(decodedLink)
		if len(matches) > 1 {
			doubanID, err = strconv.ParseInt(matches[1], 10, 64)
			if err != nil {
				d.logger.Errorf("解析豆瓣ID失败: %v", err)
				return
			}
			found = true // 标记为已找到
		}
	})

	err := c.Visit(searchURL)
	if err != nil {
		d.logger.Errorf("搜索请求失败: %v, URL: %s", err, searchURL)
		return 0, err
	}
	c.Wait()

	if !found {
		return 0, fmt.Errorf("未找到与名称 '%s' 相关的豆瓣ID", name)
	}

	return doubanID, nil
}

func (d *Douban) extraInfoById(doubanID int64) *biz.Douban {
	c := colly.NewCollector()
	doubanLink := fmt.Sprintf(DoubanDetail, doubanID)
	data := &biz.Douban{
		DoubanId:   doubanID,
		DoubanLink: doubanLink,
	}

	c.OnHTML("div#content", func(e *colly.HTMLElement) {
		data.Name = e.ChildText("h1 span[property='v:itemreviewed']")
		data.Directors = e.ChildTexts("a[rel='v:directedBy']")
		data.Actors = e.ChildTexts("a[rel='v:starring']")
		data.Genres = e.ChildTexts("span[property='v:genre']")
		data.Rating = e.ChildText("strong.rating_num")
		data.Year = regexp.MustCompile(`\d+`).FindString(e.ChildText("span.year"))
		data.Introduction = strings.TrimSpace(e.ChildText("span[property='v:summary']"))
		data.ReleaseDate = e.ChildText("span[property='v:initialReleaseDate']")
		data.EpisodeCount = e.ChildText("span:contains('集数:') + text()")
		data.EpisodeDuration = e.ChildText("span[property='v:runtime']")

		posterLink := e.ChildAttr("div#mainpic a img", "src")
		if posterLink != "" {
			data.PosterLink = posterLink
			data.PosterData = d.fetchMediaImage(posterLink)
		}

		var writers []string
		e.ForEach("span.pl", func(_ int, el *colly.HTMLElement) {
			text := strings.TrimSpace(el.Text)
			switch text {
			case "编剧":
				if siblings := el.DOM.NextAll(); siblings.Length() > 0 {
					writers = strings.Split(strings.ReplaceAll(siblings.First().Text(), " ", ""), "/")
				}
			case "集数:":
				if sibling := el.DOM.Next(); sibling.Length() > 0 {
					data.EpisodeCount = sibling.Text()
				}
			case "单集片长:":
				if sibling := el.DOM.Next(); sibling.Length() > 0 && data.EpisodeDuration == "" {
					data.EpisodeDuration = sibling.Text()
				}
			}
		})
		data.Writers = writers
	})

	c.OnError(func(r *colly.Response, err error) {
		d.logger.Errorf("详情页请求失败: %s, 错误: %v", r.Request.URL, err)
	})

	err := c.Visit(doubanLink)
	if err != nil {
		d.logger.Errorf("访问详情页失败: %v, URL: %s", err, doubanLink)
	}
	c.Wait()

	return data
}

func (d *Douban) fetchMediaImage(imageURL string) []byte {
	imgCollector := colly.NewCollector()
	var imgData []byte
	var visited bool

	imgCollector.OnResponse(func(r *colly.Response) {
		if !visited {
			imgData = r.Body
			visited = true
		}
	})

	err := imgCollector.Visit(imageURL)
	if err != nil {
		d.logger.Errorf("图片下载失败: %v, URL: %s", err, imageURL)
	}
	return imgData
}
