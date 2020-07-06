import "dart:html" as html;

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
  html.ParagraphElement p =
      html.querySelector("#RipVanWinkle") as html.ParagraphElement;
  p.text = "Wake up, sleepy head!";
}
