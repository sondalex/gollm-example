project('go-llama.cpp', 'cpp', default_options:[
    'cpp_std=c++11', 'c_std=c11', 'b_staticpic=true',
    'optimization=3'
  ]
)

gollamaheaders = include_directories('.')

# Add the warning flags to the compiler options
add_project_arguments(
  [
    '-Wall', 
    '-Wextra', 
    '-Wpedantic',
    '-Wcast-qual',
    '-Wdouble-promotion',
    '-Wshadow',
    '-Wstrict-prototypes',
    '-Wpointer-arith',
    '-Wno-unused-function'
  ], 
  language:'c'
)

add_project_arguments(
  [
    '-Wall',
    '-Wextra',
    '-Wpedantic',
    '-Wcast-qual',
    '-Wno-unused-function'
  ],
  language:'cpp'
)

cmake = import('cmake')
opt_var = cmake.subproject_options()

opt_var.set_override_option('cpp_std', 'c++11')
opt_var.set_override_option('c_std', 'c11')
opt_var.set_override_option('b_staticpic', 'true')
opt_var.set_override_option('optimization', '3')


llama_cpp = cmake.subproject('llama.cpp', options: opt_var)

root = meson.global_build_root()

llama_root = join_paths(root, 'subprojects', 'llama.cpp')

outdir = join_paths(meson.project_source_root(), 'llama.cpp')

# Get the libraries from the llama_cpp project


common_dep = llama_cpp.dependency('common')
llama_dep = llama_cpp.dependency('llama')
ggml_dep = llama_cpp.dependency('ggml')

install_dir = join_paths(meson.global_source_root(), 'lib')

source_files = [
    join_paths(llama_root, 'libllama.a.p', 'llama.cpp.o'), 
    join_paths(llama_root, 'libggml.a.p', 'ggml.c.o'), 
    join_paths(llama_root, 'libggml.a.p', 'ggml-alloc.c.o'), 
    join_paths(llama_root, 'libggml.a.p', 'k_quants.c.o'), 
    join_paths(llama_root, 'libcommon.a.p', 'common_common.cpp.o'), 
    join_paths(llama_root, 'libcommon.a.p', 'common_common.cpp.o'), 
    join_paths(llama_root, 'libcommon.a.p', 'common_grammar-parser.cpp.o')
]

copy_rename_rule = custom_target(
    'copy_and_rename_files',
    input: source_files,
    output: [
        'llama.o',
        'ggml.o',
        'ggml-alloc.o',
        'k_quants.o',
        'common_common.o',
        'grammar-parser.o'
    ],
    command: [
        'cp', 
        join_paths(llama_root, 'libllama.a.p', 'llama.cpp.o'), join_paths(outdir, 'llama.o'),
        '&&',
        'cp', 
        join_paths(llama_root, 'libggml.a.p', 'ggml.c.o'), join_paths(outdir, 'ggml.o'),
        '&&',
        'cp', 
        join_paths(llama_root, 'libggml.a.p', 'ggml-alloc.c.o'), join_paths(outdir, 'ggml-alloc.o'),
        '&&',
        'cp', 
        join_paths(llama_root, 'libggml.a.p', 'k_quants.c.o'), join_paths(outdir, 'k_quants.o'),
        '&&',
        'cp', 
        join_paths(llama_root, 'libcommon.a.p', 'common_common.cpp.o'), join_paths(outdir, 'common.o'),
        '&&',
        'cp', 
        join_paths(llama_root, 'libcommon.a.p', 'common_grammar-parser.cpp.o'), join_paths(outdir, 'grammar-parser.o')
    ],
)

go_llama_root = meson.project_build_root()

llama_cpp_root = join_paths(meson.global_build_root(), 'subprojects', 'llama.cpp')
copy_rename_rule2 = custom_target(
    'copy_and_rename_files2',
    output: [
        'libembdinput_custom.a',
        'libllama_custom.a'
    ],
    command: [
        'cp', join_paths(llama_cpp_root, 'libembdinput.a'), 
        join_paths(meson.project_source_root(), 'libembdinput.a'),
        '&&',
        'cp', join_paths(llama_cpp_root, 'libllama.a'), 
        join_paths(meson.project_source_root(), 'libllama.a')
    ],
)

