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
late html.ButtonElement TagButton;

late html.InputElement CardQuestionInput;
late html.SelectElement CardTagSelect;
late html.TextAreaElement CardAnswerText;
late html.ButtonElement CardButton;

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
    TagButton = html.querySelector("#buttonAddTags") as html.ButtonElement;
    TagButton.onClick.listen(addTags);
  }

  {
    //Create Cards.
    CardQuestionInput = html.querySelector("#cardQuestion") as html.InputElement;
    CardTagSelect = html.querySelector("#cardTag") as html.SelectElement;
    CardAnswerText = html.querySelector("#cardAnswer") as html.TextAreaElement;
    CardButton = html.querySelector("#buttonAddCards") as html.ButtonElement;
    CardButton.onClick.listen(addCards);
    // we need to call populateTags here so that html page can be populated with tags at startup
    populateTags();
  }
}

void addTags(html.Event e) {
  var t = srs.Tag(TagNameInput.value, TagDescriptionInput.value);
  AllTags.add(t);
  TagNameInput.value = "";
  TagDescriptionInput.value = "";

  populateTags();
  print("""{"event": "addTags", "Tag": $t}""");
}

void addCards(html.Event e) {
  List<srs.Tag> _tags = [];
  var _selected_tags = CardTagSelect.selectedOptions;
  _selected_tags.forEach((i) {
    _tags.add(srs.Tag(i.value, "some stuff"));
  });

  var c = srs.Card(CardQuestionInput.value, CardAnswerText.value, _tags);
  CardQuestionInput.value = "";
  CardAnswerText.value = "";

  populateTags();
  print("""{"event": "addCards", "Card": $c}""");
}

void populateTags() {
  CardTagSelect.children = [];
  CardTagSelect.size = AllTags.length + 1;
  AllTags.forEach((i) {
    var newTagOpt = html.OptionElement();
    newTagOpt.value = i.name;
    newTagOpt.text = i.name;

    //TODO: find a way to dedupe.
    CardTagSelect.children.add(newTagOpt);
  });
}
