#-------------------------------------------------------------------------------------------------------------
# author: Alex Spina
# license: Have your way with me.
# date: 12/22/2019
#-------------------------------------------------------------------------------------------------------------
FROM alexmspina/godev:v1.0.2
#
## Copy init.vim to container so plug can install vim packages
COPY ./build/godev/init.vim /root/.config/nvim/init.vim
#
## Install neovim plugins
RUN nvim +PlugInstall +qall
#
## switch to root directory
WORKDIR /root
#
## copy project directory to root/{project_name}
COPY . ./project
#
## switch to project directory
WORKDIR /root/project
#
## download dependencies from go.mod and go.sum
RUN go get golang.org/x/tools/gopls@latest && \
    go mod download
#
## start z shell
ENTRYPOINT [ "/usr/bin/fish" ]