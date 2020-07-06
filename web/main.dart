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

late html.InputElement CardQuestionInput;
late html.InputElement CardTagInput;
late html.TextAreaElement CardAnswerText;

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
    TagDescriptionInput =
        html.querySelector("#tagDescription") as html.InputElement;
    TagDescriptionInput.onChange.listen(addTags);
  }

  {
    //Create Cards.
    CardQuestionInput =
        html.querySelector("#cardQuestion") as html.InputElement;
    // TODO: card tags should be a dropdown selection.
    CardTagInput = html.querySelector("#cardTag") as html.InputElement;
    CardAnswerText = html.querySelector("#cardAnswer") as html.TextAreaElement;
    CardAnswerText.onChange.listen(addCards);
  }
}

void addTags(html.Event e) {
  var t = srs.Tag(TagNameInput.value, TagDescriptionInput.value);
  AllTags.add(t);

  print("""{"event": "addTags", "Tag": $t}""");
}

void addCards(html.Event e) {
  print(
      "card Q: ${CardQuestionInput.value} card A: ${CardAnswerText.value} card Tag: ${CardTagInput}  ");
  // var c = srs.Card("name?", "My name is Kapombe.",
  //     srs.Tag("cs", "computer science general knowledge"));
  // print("""{"event": "addCards", "Card": $c}""");
}
