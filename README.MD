# 视频 1 演示:如何使用 github 提供的 Pages 服务,让自己的文档随处可用

- 步骤:git clone https://github.com/hfbhfb/my-web
- 步骤:创建后缀为-io 的 github 项目: 例如:https://github.com/hfbhfb/my-web2-io https://github.com/hfbhfb/my-web-io
- 步骤:git remote add branchccc git@github.com:hfbhfb/my-web2.git #注意这个不带-io 后缀,只是在项目打一个标签
- 步骤:go run deploy.go #用的是 go 写的类脚本逻辑:带参数支持 cname
- 步骤:在 github 项目设置项目 pages 服务(https://github.com/hfbhfb/my-web2-io/settings)

# 视频 2 演示:基本使用

- 修改个人简历
- 在 golang 主题下增加两个文档
- 增加目录并放到首页

# 项目感谢: vuepress theme-vdoing, 自己在视频中学了很多东西,自己也为平台贡献自己的力量
