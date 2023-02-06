# Release steps for emco-gui on GitLab

This document outlines the release process for emco-gui on GitLab.
The version as of this writing is `22.03`. Please modify to the correct version according to the release undertaken.

## Tagging the source code

    git tag v22.03

## Pushing container images

    EMCO_VERSION=22.03

    cd emco-gui # this is the git repo dir
    docker login registry.gitlab.com

    # emco-gui
    docker build -t emco-gui:latest .
    docker tag emco-gui:latest emco-gui:$EMCO_VERSION
    docker tag emco-gui:$EMCO_VERSION registry.gitlab.com/project-emco/ui/emco-gui/emco-gui:$EMCO_VERSION
    docker tag emco-gui:latest registry.gitlab.com/project-emco/ui/emco-gui/emco-gui:latest
    docker push registry.gitlab.com/project-emco/ui/emco-gui/emco-gui:latest
    docker push registry.gitlab.com/project-emco/ui/emco-gui/emco-gui:$EMCO_VERSION

    # emco-gui-dbhook
    docker build -t emco-gui-dbhook:latest db_udpate
    docker tag emco-gui-dbhook:latest emco-gui-dbhook:$EMCO_VERSION
    docker tag emco-gui-dbhook:$EMCO_VERSION registry.gitlab.com/project-emco/ui/emco-gui/emco-gui-dbhook:$EMCO_VERSION
    docker tag emco-gui-dbhook:latest registry.gitlab.com/project-emco/ui/emco-gui/emco-gui-dbhook:latest
    docker push registry.gitlab.com/project-emco/ui/emco-gui/emco-gui-dbhook:latest
    docker push registry.gitlab.com/project-emco/ui/emco-gui/emco-gui-dbhook:$EMCO_VERSION

    # emco-gui-authgw
    docker build -t emco-gui-authgw:latest authgateway
    docker tag emco-gui-authgw:latest emco-gui-authgw:$EMCO_VERSION
    docker tag emco-gui-authgw:$EMCO_VERSION registry.gitlab.com/project-emco/ui/emco-gui/emco-gui-authgw:$EMCO_VERSION
    docker tag emco-gui-authgw:latest registry.gitlab.com/project-emco/ui/emco-gui/emco-gui-authgw:latest
    docker push registry.gitlab.com/project-emco/ui/emco-gui/emco-gui-authgw:latest
    docker push registry.gitlab.com/project-emco/ui/emco-gui/emco-gui-authgw:$EMCO_VERSION

    # emco-gui-middleend
    docker build -t emco-gui-middleend:latest guimiddleend
    docker tag emco-gui-middleend:latest emco-gui-middleend:$EMCO_VERSION
    docker tag emco-gui-middleend:$EMCO_VERSION registry.gitlab.com/project-emco/ui/emco-gui/emco-gui-middleend:$EMCO_VERSION
    docker tag emco-gui-middleend:latest registry.gitlab.com/project-emco/ui/emco-gui/emco-gui-middleend:latest
    docker push registry.gitlab.com/project-emco/ui/emco-gui/emco-gui-middleend:latest
    docker push registry.gitlab.com/project-emco/ui/emco-gui/emco-gui-middleend:$EMCO_VERSION

## Creating release page and tarballs

Go to the [GitLab Releases page](https://gitlab.com/project-emco/ui/emco-gui/-/releases), and click *New release*.
Choose the tag name `v22.03` created earlier. Name the release title simply as `22.03`. Choose the corresponding milestone, i.e. 22.03. Finally, write the release notes, add any extra release assets, and click *Create release*.
