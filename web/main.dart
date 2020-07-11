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
late html.ButtonElement TagButton;

late html.InputElement CardQuestionInput;
late html.SelectElement AddCardTagSelect;
late html.TextAreaElement CardAnswerText;
late html.ButtonElement CardButton;

late html.DivElement ReviewCardTagDiv;
late html.ButtonElement ReviewCardsButton;

late html.ParagraphElement CurrentQuestionP;
late html.ButtonElement ShowCardAnswerButton;

var AllTags = srs.generateDefaultTags();
List<srs.Card> AllCards = [];

List<srs.Card> Cards2Review = [];

void main() {
  MainDiv = html.querySelector("#app") as html.DivElement;

  {
    // Create Tags.
    TagNameInput = html.querySelector("#tagName") as html.InputElement;
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
    //select cards for review.
    ReviewCardTagDiv = html.querySelector("#reviewCardTags") as html.DivElement;
    ReviewCardsButton = html.querySelector("#buttonReviewCards") as html.ButtonElement;
    ReviewCardsButton.onClick.listen(reviewCards);
  }

  {
    // show card for review
    CurrentQuestionP = html.querySelector("#currentQuestion") as html.ParagraphElement;
    ShowCardAnswerButton = html.querySelector("#buttonShowCardAnswer") as html.ButtonElement;
    ShowCardAnswerButton.onClick.listen(renderCardsForReview);
  }

  // we need to call populateTags here so that html page can be populated with tags at startup
  populateTags();
}

void addTags(html.Event e) {
  var t = srs.Tag(TagNameInput.value);
  AllTags.add(t);
  TagNameInput.value = "";

  populateTags();
  print("""{"event": "addTags", "Tag": $t}""");
}

void addCards(html.Event e) {
  List<srs.Tag> _tags = [];
  var _selected_tags = AddCardTagSelect.selectedOptions;
  _selected_tags.forEach((i) {
    _tags.add(srs.Tag(i.value));
  });

  var c = srs.Card(CardQuestionInput.value, CardAnswerText.value, _tags);
  AllCards.add(c);
  CardQuestionInput.value = "";
  CardAnswerText.value = "";

  populateTags();
  print("""{"event": "addCards", "Card": $c}""");
}

void reviewCards(html.Event e) {
  List<String> _selectedTags = [];
  ReviewCardTagDiv.children.forEach((el) {
    if (el is html.CheckboxInputElement) {
      if (el.checked) {
        _selectedTags.add(el.value);
      }
    }
  });

  _selectedTags.forEach((el) {
    print("selected Tag: ${el}");
  });

  AllCards.forEach((aC) {
    // TODO: this is too much complexity, simplify.
    _selectedTags.forEach((st) {
      aC.tags.forEach((acT) {
        if (acT.name == st) {
          Cards2Review.add(aC);
        }
      });
    });
  });

  populateTags();
  print("""{"event": "reviewCards" "Cards2Review": $Cards2Review}""");
}

void renderCardsForReview(html.Event e) {
  Cards2Review.forEach((c) {
    CurrentQuestionP.text = c.question;
  });

  /////////
  // CurrentQuestionP =
  //     html.querySelector("#currentQuestion") as html.ParagraphElement;
  // ShowCardAnswerButton =
  //     html.querySelector("#buttonShowCardAnswer") as html.ButtonElement;

  print("""{"event": "renderCardsForReview" "Cards2Review": $Cards2Review}""");
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
