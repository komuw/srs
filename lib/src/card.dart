import "dart:math" as math;
import "./sm.dart" as sm;
import "./tag.dart" as tag;

class Card {
  String question;
  String answer;

  // all the datetimes should be in UTC
  late DateTime createdAt;
  late DateTime updatedAt;
  late DateTime nextReviewDate;

  int maxRatings = 50; // only keep the last 50 ratings
  late List<sm.Rating> historyOfRatings;
  late Set<tag.Tag> tags;

  Card(this.question, this.answer, tag.Tag t) {
    var now = DateTime.now().toUtc();

    createdAt = now;
    updatedAt = now;
    nextReviewDate = now.add(Duration(days: 1));

    historyOfRatings = [sm.Rating.Hard];

    tags = tag.generateDefaultTags();
    tags.add(t);
  }

  @override
  String toString() {
    var _qSubst = question.substring(0, math.min(6, question.length));
    var _aSubst = answer.substring(0, math.min(6, answer.length));

    return "Card(Q:$_qSubst, A:$_aSubst, next:${nextReviewDate.day}-${nextReviewDate.month}-${nextReviewDate.year})";
  }

  void update(sm.Rating r) {
    historyOfRatings.add(r);
    var days2NextReview = sm.sm2(historyOfRatings);
    nextReviewDate = nextReviewDate.add(Duration(days: days2NextReview as int));

    var numRatings = historyOfRatings.length;
    if (numRatings > maxRatings) {
      historyOfRatings = historyOfRatings.sublist(numRatings - maxRatings);
    }
  }
}
