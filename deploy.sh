#!/usr/bin/env sh

for url in $(git remote -v)
do
filename=${url%.*c}
    echo $url| awk -F'git@github.com:' '{print $0,$1,$2}'
done


echo $CODING_TOKEN
exit

# 确保脚本抛出遇到的错误
set -e

# 生成静态文件
npm run build

# 进入生成的文件夹
cd docs/.vuepress/dist

# deploy to github
echo 'b.hfbhfb.com' > CNAME
if [ -z "$GITHUB_TOKEN" ]; then
  msg='deploy'
  githubUrl=git@github.com:hfbhfb/for-open.git
else
  msg='来自github actions的自动部署'
  githubUrl=https://hfbhfb:${GITHUB_TOKEN}@github.com/hfbhfb/for-open.git
  git config --global user.name "hefabao"
  git config --global user.email "hefabao@126.com"
fi
git init
git add -A
git commit -m "${msg}"
git push -f $githubUrl master:gh-pages # 推送到github

# deploy to coding
echo 'www.hfbhfb.com\nhfbhfb.com' > CNAME  # 自定义域名
if [ -z "$CODING_TOKEN" ]; then  # -z 字符串 长度为0则为true；$CODING_TOKEN来自于github仓库`Settings/Secrets`设置的私密环境变量
  codingUrl=git@e.coding.net:xgy/xgy.git
else
  codingUrl=https://HmuzsGrGQX:${CODING_TOKEN}@e.coding.net/xgy/xgy.git
fi
git add -A
git commit -m "${msg}"
git push -f $codingUrl master # 推送到coding

cd - # 退回开始所在目录
rm -rf docs/.vuepress/dist