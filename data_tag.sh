function push_tag() {
    tag=$(date "+%y.%m.%d.%s")
    echo $tag

    git tag $tag
    # git tag -a $tag -m ""
    #推送单个标签到远端
    git push git $tag
    git push git --all
}
git commit -am "$1"
push_tag
