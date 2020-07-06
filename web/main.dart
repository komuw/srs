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
  html.DivElement div = html.querySelector("#app") as html.DivElement;
  print("p.type: ${div.runtimeType}");
  div.text = "New app text, via Dartlang!";
}
