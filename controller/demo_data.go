package controller

//var DemoVideos = []*proto.Video{
//	{
//		Id:            1,
//		Author:        DemoUser,
//		PlayUrl:       "http://vfx.mtime.cn/Video/2019/02/04/mp4/190204084208765161.mp4",
//		CoverUrl:      "http://vfx.mtime.cn/Video/2019/02/04/mp4/190204084208765161.mp4?x-oss-process=video/snapshot,t_30000,f_jpg,w_0,h_0,m_fast,ar_auto",
//		FavoriteCount: 0,
//		CommentCount:  0,
//		IsFavorite:    true,
//		Title:         "咦哈",
//	}, {
//		Id:            2,
//		Author:        DemoUser,
//		PlayUrl:       "http://vfx.mtime.cn/Video/2019/03/21/mp4/190321153853126488.mp4",
//		CoverUrl:      "http://vfx.mtime.cn/Video/2019/03/21/mp4/190321153853126488.mp4?x-oss-process=video/snapshot,t_30000,f_jpg,w_0,h_0,m_fast,ar_auto",
//		FavoriteCount: 0,
//		CommentCount:  0,
//		IsFavorite:    true,
//		Title:         "我嫩亲爹",
//	}, {
//		Id:            3,
//		Author:        DemoUser,
//		PlayUrl:       "https://v-cdn.zjol.com.cn/276990.mp4",
//		CoverUrl:      "https://v-cdn.zjol.com.cn/276990.mp4?x-oss-process=video/snapshot,t_30000,f_jpg,w_0,h_0,m_fast,ar_auto",
//		FavoriteCount: 0,
//		CommentCount:  0,
//		IsFavorite:    true,
//		Title:         "我嫩爷",
//	}, {
//		Id:            4,
//		Author:        DemoUser,
//		PlayUrl:       "http://vjs.zencdn.net/v/oceans.mp4",
//		CoverUrl:      "http://vjs.zencdn.net/v/oceans.mp4?x-oss-process=video/snapshot,t_30000,f_jpg,w_0,h_0,m_fast,ar_auto",
//		FavoriteCount: 0,
//		CommentCount:  0,
//		IsFavorite:    false,
//		Title:         "我嫩亲爷",
//	},
//}

//var DemoComments = []Comment{
//	{
//		Id:         1,
//		User:       *DemoUser,
//		Content:    "我是你爸爸",
//		CreateDate: "05-01",
//	},
//}

var DemoUser = &TestUser{
	Id:            1,
	Name:          "逯振宇",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}
