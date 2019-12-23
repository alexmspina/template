#-------------------------------------------------------------------------------------------------------------
# author: Alex Spina
# license: Have your way with me.
# date: 10/31/2019
#-------------------------------------------------------------------------------------------------------------
FROM alexmspina/certdev:v1.0.0
#
## Copy init.vim to container so plug can install vim packages
COPY ./build/certdev/init.vim /root/.config/nvim/init.vim
#
## switch to root directory
WORKDIR /root
#
## start z shell
ENTRYPOINT [ "/usr/bin/fish" ]