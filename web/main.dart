import "dart:html" as html;
import "package:srs/srs.dart" as srs;

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

late html.DivElement MainDiv;
late html.InputElement TagNameInput;
late html.InputElement TagDescriptionInput;

var AllTags = srs.generateDefaultTags();

void main() {
  MainDiv = html.querySelector("#app") as html.DivElement;

  // html.ButtonElement AddTagsButton =
  //     html.querySelector("#buttonAddTags") as html.ButtonElement;
  // print("AddTagsButton.type: ${AddTagsButton.runtimeType}");
  // AddTagsButton.onClick.listen(addTags);

  {
    // Create Tags.
    TagNameInput = html.querySelector("#tagName") as html.InputElement;
    TagDescriptionInput = html.querySelector("#tagDescription") as html.InputElement;
    TagDescriptionInput.onChange.listen(addTags);
  }
}

void addTags(html.Event e) {
  var t = srs.Tag(TagNameInput.value, TagDescriptionInput.value);
  AllTags.add(t);

  print("""{"event": "addTags", "Tag": $t}""");
}
