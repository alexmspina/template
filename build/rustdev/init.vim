" Plugins will be downloaded under the specified directory.
call plug#begin('/root/.config/nvim/plugged')

" Declare the list of plugins.
Plug 'scrooloose/nerdTree'
Plug 'mattn/emmet-vim'
Plug 'tpope/vim-surround'
Plug 'ctrlpvim/ctrlp.vim'
Plug 'bling/vim-airline'
Plug 'tpope/vim-eunuch'
Plug 'dag/vim-fish'
Plug 'elzr/vim-json'
Plug 'ekalinin/dockerfile.vim'

" List ends here. Plugins become visible to Vim after this call.
call plug#end()

" key mappings
nmap <C-n> :NERDTreeToggle<CR>

" key bindings
let g:user_emmet_expandabbr_key = '<C-a>,'
