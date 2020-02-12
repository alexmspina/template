#-------------------------------------------------------------------------------------------------------------
# author: Alex Spina
# license: ...
# date: 01/04/2019
#-------------------------------------------------------------------------------------------------------------
FROM alexmspina/rustdev:v0.1.0
#
## Copy init.vim to container so plug can install vim packages
COPY ./build/rustdev/init.vim /root/.config/nvim/init.vim
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
## start z shell
ENTRYPOINT [ "/usr/bin/fish" ]
