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

void main() {
  var defaultTags = srs.generateDefaultTags();
  print("defaultTags: ${defaultTags}");
  MainDiv = html.querySelector("#app") as html.DivElement;

  // html.ButtonElement AddTagsButton =
  //     html.querySelector("#buttonAddTags") as html.ButtonElement;
  // print("AddTagsButton.type: ${AddTagsButton.runtimeType}");
  // AddTagsButton.onClick.listen(addTags);

  TagNameInput = html.querySelector("#tagName") as html.InputElement;
  TagDescriptionInput =
      html.querySelector("#tagDescription") as html.InputElement;

  TagDescriptionInput.onChange.listen(addTags);
}

void addTags(html.Event e) {
  print("addTags func called.");
  print("with event: $e");

  var tName = TagNameInput.value;
  var tDescription = TagDescriptionInput.value;

  print("tName: $tName , tDescription: $tDescription");
  var t = srs.Tag(tName, tDescription);

  print("new Tag created: $t");
}
