# The incantation of flags and env variables needed to support NNBD
# may change(in form, name & whatever) in the future.
# Also, once NNBD is released in a stable version, we'll no longer need the incantations
#
# 1. https://github.com/dart-lang/sdk/issues/42465
# 2. https://github.com/dart-lang/sdk/issues/41853
# 3. https://github.com/dart-lang/pub/pull/2542


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

serve:
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
        printf "\n run webdev serve ::\n" && \
        webdev serve \
                --debug \
                --debug-extension \
                --injected-client \
                --log-requests \
                --output=web:output \
                web:8080
        # --release


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
