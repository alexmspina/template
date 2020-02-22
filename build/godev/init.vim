" Plugins will be downloaded under the specified directory.
call plug#begin('/root/.config/nvim/plugged')

" Declare the list of plugins.
Plug 'scrooloose/nerdTree'
Plug 'tpope/vim-surround'
Plug 'ctrlpvim/ctrlp.vim'
Plug 'bling/vim-airline'
Plug 'dag/vim-fish'
Plug 'tpope/vim-eunuch'
Plug 'elzr/vim-json'
Plug 'ekalinin/dockerfile.vim'
Plug 'terryma/vim-multiple-cursors'
Plug 'tpope/vim-commentary'
Plug 'rstacruz/vim-closer'
Plug 'fatih/vim-go', { 'do': ':GoUpdateBinaries' }

" List ends here. Plugins become visible to Vim after this call.
call plug#end()

" set line numbers
set number

" map leader key
let mapleader = ";"

" show hidden files in NERDTree by default
let NERDTreeShowHidden=1

" NERDTree key mappings
nmap <Leader>nt :NERDTreeToggle<CR>

" multiple cursor key mappings
let g:multi_cursor_use_default_mapping=0
let g:multi_cursor_start_word_key = '<C-n>'
let g:multi_cursor_select_all_word_key = '<A-n>'
let g:multi_cursor_start_key = 'g<C-n>'
let g:multi_cursor_select_all_key = 'g<A-n>'
let g:multi_cursor_next_key = '<C-n>'
let g:multi_cursor_prev_key = '<C-p>'
let g:multi_cursor_skip_key = '<C-x>'
let g:multi_cursor_quit_key = '<Esc>'

" set vim-go to auto-import and fmt on save
let g:go_fmt_command = "goimports"
