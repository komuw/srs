import "dart:html";

/*
webdev serve \
        --debug \
        --debug-extension \
        --injected-client \
        --log-requests \
        --output=web:output \
        web:8080
# --release
*/
void main() {
  querySelector("#RipVanWinkle")?.text = "Wake up, sleepy head!";
}
