get:
	@export DART_VM_OPTIONS='--enable-experiment=non-nullable'
	@printf "\n dart format::\n" && dartfmt --overwrite --follow-links --line-length=100 --fix .
	@printf "\n dartanalyzer::\n" && dartanalyzer \
        --options=analysis_options.yaml \
        --no-declaration-casts \
        --no-implicit-casts \
        --no-implicit-dynamic \
        --fatal-warnings \
        --enable-experiment=non-nullable .
	@printf "\n pub get::\n" && pub get


run:
	@export DART_VM_OPTIONS='--enable-experiment=non-nullable'
	@printf "\n dart format::\n" && dartfmt --overwrite --follow-links --line-length=100 --fix .
	@printf "\n dartanalyzer::\n" && dartanalyzer \
        --options=analysis_options.yaml \
        --no-declaration-casts \
        --no-implicit-casts \
        --no-implicit-dynamic \
        --fatal-warnings \
        --enable-experiment=non-nullable .
	@printf "\n run dart::\n" && dart --enable-experiment=non-nullable lib/tag.dart

# you can run single testcase as;
# dart format .; pub run test .
t:
	@export DART_VM_OPTIONS='--enable-experiment=non-nullable'
	@printf "\n dart format::\n" && dartfmt --overwrite --follow-links --line-length=100 --fix .
	@printf "\n dartanalyzer::\n" && dartanalyzer \
        --options=analysis_options.yaml \
        --no-declaration-casts \
        --no-implicit-casts \
        --no-implicit-dynamic \
        --fatal-warnings \
        --enable-experiment=non-nullable .
	@printf "\n run tests::\n" && pub run --enable-experiment=non-nullable test .

