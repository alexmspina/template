" Plugins will be downloaded under the specified directory.
call plug#begin('/root/.config/nvim/plugged')

" Declare the list of plugins.
Plug 'scrooloose/nerdTree'
Plug 'mattn/emmet-vim'
Plug 'tpope/vim-surround'
Plug 'ctrlpvim/ctrlp.vim'
Plug 'bling/vim-airline'
Plug 'dag/vim-fish'
Plug 'tpope/vim-eunuch'
Plug 'dense-analysis/ale'
Plug 'elzr/vim-json'
Plug 'ekalinin/dockerfile.vim'
Plug 'maxmellon/vim-jsx-pretty'
Plug 'yuezk/vim-js'

" List ends here. Plugins become visible to Vim after this call.
call plug#end()

" key mappings
nmap <C-n> :NERDTreeToggle<CR>

" key bindings
let g:user_emmet_expandabbr_key = '<C-a>,'

" standard js configs
let g:ale_fixers = {}
let g:ale_linters = {
			\'javascript': ['standard'],
			\}
let g:ale_fixers = {
			\'javascript': ['standard'],
			\}
let g:ale_lint_on_save = 1
let g:ale_fix_on_save = 1
let g:ale_javascript_standard_use_global = 1
