MartianBlog功能架构
=================

前台
---

前台共有两个页面 —— `首页`，`文章详情`

1. 首页

	首页显示所有文章的完整列表。从上到下依次是：
	
	+ logo
	+ title
	+ subtitle
	+ article
		+ title
		+ date + tags
		+ content
		+ author
	+ footer
	
2. 文章详情

	文章详情页显示选定文章的详细内容，已经评论模块。从上到下依次是：
	
	+ logo
	+ title
	+ subtitle
	+ article
		+ title
		+ date + tags
		+ content
		+ author
	+ discussion
	+ footer
	

后台
---

后台共五个页面 —— `登录`，`控制台`，`文章列表`，`文章编辑`，`个人信息`

1. 登录
	
	用邮箱和密码来登录，登录成功后跳转到控制台页面。
	
2. 控制台
	
	放置**Google Analytics**的统计图表。具体功能待定。
	
3. 文章列表