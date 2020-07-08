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
late html.SelectElement AddCardTagSelect;
late html.TextAreaElement CardAnswerText;
late html.ButtonElement CardButton;

late html.DivElement ReviewCardTagDiv;
late html.ButtonElement ReviewCardsButton;

var AllTags = srs.generateDefaultTags();

void main() {
  MainDiv = html.querySelector("#app") as html.DivElement;

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
    AddCardTagSelect = html.querySelector("#addCardTags") as html.SelectElement;
    CardAnswerText = html.querySelector("#cardAnswer") as html.TextAreaElement;
    CardButton = html.querySelector("#buttonAddCards") as html.ButtonElement;
    CardButton.onClick.listen(addCards);
  }

  {
    //Review Cards.
    ReviewCardTagDiv = html.querySelector("#reviewCardTags") as html.DivElement;
    ReviewCardsButton = html.querySelector("#buttonReviewCards") as html.ButtonElement;
    ReviewCardsButton.onClick.listen(reviewCards);
  }

  // we need to call populateTags here so that html page can be populated with tags at startup
  populateTags();
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
  var _selected_tags = AddCardTagSelect.selectedOptions;
  _selected_tags.forEach((i) {
    _tags.add(srs.Tag(i.value, "some stuff"));
  });

  var c = srs.Card(CardQuestionInput.value, CardAnswerText.value, _tags);
  CardQuestionInput.value = "";
  CardAnswerText.value = "";

  populateTags();
  print("""{"event": "addCards", "Card": $c}""");
}

void reviewCards(html.Event e) {
  List<html.CheckboxInputElement> _selectedTags = [];
  ReviewCardTagDiv.children.forEach((el) {
    if (el is html.CheckboxInputElement) {
      if (el.checked) {
        _selectedTags.add(el);
      }
    }
  });

  _selectedTags.forEach((el) {
    print("selected Tag: ${el.value}");
  });

  populateTags();
  print("""{"event": "reviewCards"}""");
}

void populateTags() {
  AddCardTagSelect.children = [];
  AddCardTagSelect.size = AllTags.length + 1;
  AllTags.forEach((i) {
    var newTagOpt = html.OptionElement();
    newTagOpt.value = i.name;
    newTagOpt.text = i.name;
    AddCardTagSelect.children.add(newTagOpt);
  });

  ReviewCardTagDiv.children = [];
  AllTags.forEach((i) {
    var newCheckbox = html.CheckboxInputElement();
    newCheckbox.name = i.name;
    newCheckbox.text = i.name;
    newCheckbox.value = i.name;
    newCheckbox.id = i.name;

    var label = html.LabelElement();
    label.htmlFor = i.name;
    label.text = i.name;

    ReviewCardTagDiv.children.add(newCheckbox);
    ReviewCardTagDiv.children.add(label);
    ReviewCardTagDiv.children.add(html.BRElement());
  });
}
