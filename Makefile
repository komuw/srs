get:
	@export DART_VM_OPTIONS='--enable-experiment=non-nullable, --no-null-safety' && \
        printf "\n dart format::\n" && \
        dartfmt --overwrite --follow-links --line-length=100 --fix . && \
        printf "\n run dartanalyzer::\n" && \
        dartanalyzer --enable-experiment=non-nullable \
        --options=analysis_options.yaml \
        --no-declaration-casts \
        --no-implicit-casts \
        --no-implicit-dynamic \
        --fatal-warnings . && \
        printf "\n run pub get::\n" && \
        pub get


run:
	@export DART_VM_OPTIONS='--enable-experiment=non-nullable, --no-null-safety' && \
        printf "\n dart format::\n" && \
        dartfmt --overwrite --follow-links --line-length=100 --fix . && \
        printf "\n run dartanalyzer::\n" && \
        dartanalyzer --enable-experiment=non-nullable \
        --options=analysis_options.yaml \
        --no-declaration-casts \
        --no-implicit-casts \
        --no-implicit-dynamic \
        --fatal-warnings . && \
        printf "\n run dart::\n" && \
        dart --enable-experiment=non-nullable lib/src/tag.dart

t:
	@export DART_VM_OPTIONS='--enable-experiment=non-nullable, --no-null-safety' && \
        printf "\n dart format::\n" && \
        dartfmt --overwrite --follow-links --line-length=100 --fix . \
        printf "\n run dartanalyzer::\n" && \
        dartanalyzer --enable-experiment=non-nullable \
        --options=analysis_options.yaml \
        --no-declaration-casts \
        --no-implicit-casts \
        --no-implicit-dynamic \
        --fatal-warnings . && \
        printf "\n run tests::\n" && \
        pub run test .

