package service

import (
	pb "YYeTsBot-Go/api/yyets/v1"
	"YYeTsBot-Go/internal/biz"
	"context"
)

type ResourceService struct {
	pb.UnimplementedResourceServer
	resourceUseCase *biz.ResourceUsecase
	commentUseCase  *biz.CommentUsecase
	userUserCase    biz.UserUseCase
}

func NewResourceService(resourceUseCase *biz.ResourceUsecase, commentUseCase *biz.CommentUsecase, userUseCase *biz.UserUseCase) *ResourceService {
	return &ResourceService{resourceUseCase: resourceUseCase, commentUseCase: commentUseCase, userUserCase: *userUseCase}
}

func (s *ResourceService) CreateResource(ctx context.Context, req *pb.CreateResourceRequest) (*pb.CreateResourceReply, error) {
	return &pb.CreateResourceReply{}, nil
}
func (s *ResourceService) UpdateResource(ctx context.Context, req *pb.UpdateResourceRequest) (*pb.UpdateResourceReply, error) {
	return &pb.UpdateResourceReply{}, nil
}
func (s *ResourceService) DeleteResource(ctx context.Context, req *pb.DeleteResourceRequest) (*pb.DeleteResourceReply, error) {
	return &pb.DeleteResourceReply{}, nil
}
func (s *ResourceService) GetResource(ctx context.Context, req *pb.GetResourceRequest) (*pb.GetResourceReply, error) {
	resource, err := s.resourceUseCase.FindByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	reply := &pb.GetResourceReply{}
	reply.Id = resource.Id
	reply.Cnname = resource.CnName
	reply.Enname = resource.EnName
	reply.Aliasname = resource.AliasName
	reply.Area = resource.Area
	reply.Views = resource.Views

	dataInfo := &pb.ResourceDataInfo{
		Id:        resource.Data.Info.ID,
		Cnname:    resource.Data.Info.CnName,
		Enname:    resource.Data.Info.EnName,
		Aliasname: resource.Data.Info.AliasName,
		Channel:   resource.Data.Info.Channel,
		ChannelCn: resource.Data.Info.ChannelCN,
		Area:      resource.Data.Info.Area,
		ShowType:  resource.Data.Info.ShowType,
		Expire:    resource.Data.Info.Expire,
		Views:     resource.Data.Info.Views,
		Year:      resource.Data.Info.Year,
	}

	dataList := make([]*pb.ResourceSeason, 0)
	for _, season := range resource.Data.List {
		seasonFileItems := make(map[string]*pb.ResourceSeasonList)
		for key, items := range season.Items {
			seasonFileItems[key] = &pb.ResourceSeasonList{}
			for _, item := range items {
				files := make([]*pb.ResourceFile, 0)
				for _, seasonFile := range item.Files {
					files = append(files, &pb.ResourceFile{
						Way:     seasonFile.Way,
						WayCn:   seasonFile.WayCn,
						Address: seasonFile.Address,
						Passwd:  seasonFile.Passwd,
					})
				}
				seasonFileItems[key].Items = append(seasonFileItems[key].Items, &pb.ResourceSeasonItem{
					ItemId:     item.ItemId,
					Episode:    item.Episode,
					Name:       item.Name,
					Size:       item.Size,
					YyetsTrans: item.YyetsTrans,
					Dateline:   item.Dateline,
					Files:      files,
				})
			}
		}
		seasonItem := &pb.ResourceSeason{
			SeasonNum: season.SeasonNum,
			SeasonCn:  season.SeasonCn,
			Formats:   season.Formats,
			Items:     seasonFileItems,
		}
		dataList = append(dataList, seasonItem)
	}

	reply.Data = &pb.ResourceData{
		Info: dataInfo,
		List: dataList,
	}
	return reply, nil
}
func (s *ResourceService) ListResource(ctx context.Context, req *pb.ListResourceRequest) (*pb.ListResourceReply, error) {
	return &pb.ListResourceReply{}, nil
}
func (s *ResourceService) TopResource(ctx context.Context, req *pb.TopResourceRequest) (*pb.TopResourceReply, error) {
	resourceMap, err := s.resourceUseCase.TopResourceMap(ctx)
	if err != nil {
		return nil, err
	}

	reply := &pb.TopResourceReply{}
	for area, resources := range resourceMap {
		for _, item := range *resources {
			resource := &pb.ResourceItem{
				Data: &pb.ResourceData{
					Info: &pb.ResourceDataInfo{
						Id:        item.Data.Info.ID,
						Cnname:    item.Data.Info.CnName,
						Enname:    item.Data.Info.EnName,
						Aliasname: item.Data.Info.AliasName,
						Channel:   item.Data.Info.Channel,
						ChannelCn: item.Data.Info.ChannelCN,
						Area:      item.Data.Info.Area,
						ShowType:  item.Data.Info.ShowType,
						Expire:    item.Data.Info.Expire,
						Views:     item.Data.Info.Views,
						Year:      item.Data.Info.Year,
					},
				}}

			switch area {
			case "ALL":
				reply.ALL = append(reply.ALL, resource)
			case "US":
				reply.US = append(reply.US, resource)
			case "JP":
				reply.JP = append(reply.JP, resource)
			case "KR":
				reply.KR = append(reply.KR, resource)
			case "UK":
				reply.UK = append(reply.UK, resource)
			}
		}
	}

	reply.Class = &pb.ResourceClass{
		ALL: map[string]string{
			"$regex": ".*",
		},
		US: "美国",
		JP: "日本",
		KR: "韩国",
		UK: "英国",
	}
	return reply, nil
}

func (s *ResourceService) SearchResource(ctx context.Context, req *pb.SearchResourceRequest) (*pb.SearchResourceReply, error) {
	resources, err := s.resourceUseCase.Search(ctx, req.Keyword)
	if err != nil {
		return nil, err
	}

	comments, err := s.commentUseCase.Search(ctx, req.Keyword)
	if err != nil {
		return nil, err
	}

	reply := &pb.SearchResourceReply{}

	for _, resource := range resources {
		item := &pb.ResourceDataInfo{
			Id:        resource.Data.Info.ID,
			Cnname:    resource.Data.Info.CnName,
			Enname:    resource.Data.Info.EnName,
			Aliasname: resource.Data.Info.AliasName,
			Channel:   resource.Data.Info.Channel,
			ChannelCn: resource.Data.Info.ChannelCN,
			Area:      resource.Data.Info.Area,
			ShowType:  resource.Data.Info.ShowType,
			Expire:    resource.Data.Info.Expire,
			Views:     resource.Data.Info.Views,
			Year:      resource.Data.Info.Year,
		}
		reply.Resource = append(reply.Resource, item)
	}

	userNames, resourceIds := make([]string, 0), make([]int64, 0)
	for _, comment := range comments {
		userNames = append(userNames, comment.Username)
		resourceIds = append(resourceIds, comment.ResourceID)
	}
	userMap, err := s.userUserCase.UserGroup(ctx, userNames)
	if err != nil {
		return nil, err
	}
	resourceMap, err := s.resourceUseCase.FetchMapByIds(ctx, resourceIds)
	if err != nil {
		return nil, err
	}

	const commentOrigin string = "yyets"
	for _, comment := range comments {
		commentResource, ok := resourceMap[comment.ResourceID]
		if !ok {
			break
		}

		user := userMap[comment.Username]

		item := &pb.SearchComment{
			Username:     comment.Username,
			Date:         comment.Date,
			Comment:      comment.Content,
			CommentID:    comment.ID.Hex(),
			ResourceID:   comment.ResourceID,
			ResourceName: commentResource.Data.Info.CnName,
			Origin:       commentOrigin,
			HasAvatar:    user.Avatar != nil,
			Hash:         user.Hash,
		}
		reply.Comment = append(reply.Comment, item)
	}

	return reply, nil
}
