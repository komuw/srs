import "dart:math" as math;
import "./sm.dart" as sm;
import "./tag.dart" as tag;
import "package:intl/intl.dart" as intl;

class Card {
  final String question;
  final String answer;

  // all the datetimes should be in UTC
  late final DateTime createdAt;
  late final DateTime updatedAt;
  late DateTime nextReviewDate;

  static final int maxRatings = 50; // only keep the last 50 ratings
  late final List<sm.Rating> historyOfRatings;
  late final Set<tag.Tag> tags;

  Card(this.question, this.answer, tag.Tag t) {
    var now = getNow();

    createdAt = now;
    updatedAt = now;
    nextReviewDate = now.add(Duration(days: 1));

    historyOfRatings = [sm.Rating.Hard];

    tags = tag.generateDefaultTags();
    tags.add(t);
  }

  /// We only care about dates only without time part.
  /// If a card is due for revision today, we should be able to review it at any time of the day.
  DateTime getNow() {
    var now = DateTime.now().toUtc();
    // date only with no time component
    var dateFmt = intl.DateFormat("yyyy-MM-dd");
    var formattedNow = dateFmt.format(now);
    var _formattedNow = formattedNow + " 00:00:00.000Z";
    return DateTime.parse(_formattedNow).toUtc();
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
    nextReviewDate = nextReviewDate.add(Duration(days: days2NextReview.toInt()));

    var numRatings = historyOfRatings.length;
    if (numRatings > maxRatings) {
      historyOfRatings = historyOfRatings.sublist(numRatings - maxRatings);
    }
  }
}
