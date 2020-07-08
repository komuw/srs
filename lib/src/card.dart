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

  static final int maxRatings = 20; // only keep the last 50 ratings
  static final int longestNextReviewInterval =
      365 * 5; // card cant go longer than 5 years without review.

  List<sm.Rating> historyOfRatings = [];
  late final Set<tag.Tag> tags;

  Card(this.question, this.answer, List<tag.Tag> t) {
    var now = getNow();

    createdAt = now;
    updatedAt = now;
    nextReviewDate = now.add(Duration(days: 1));

    tags = tag.generateDefaultTags();
    tags.addAll(t);
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

    var _tagSting = "";
    tags.forEach((i) {
      _tagSting = _tagSting + i.name + ",";
    });
    return "Card(Q:$_qSubst, A:$_aSubst, next:${nextReviewDate.day}-${nextReviewDate.month}-${nextReviewDate.year}, tags:$_tagSting)";
  }

  void update(sm.Rating r) {
    historyOfRatings.add(r);
    var days2NextReview = sm.sm2(historyOfRatings);
    var days2NextReviewInt = days2NextReview.toInt();

    if (days2NextReviewInt > longestNextReviewInterval) {
      // we do not want to go long before reviewing cards.
      // also prevents, `DateTime is outside valid range` error.
      days2NextReviewInt = longestNextReviewInterval;
    }
    nextReviewDate = nextReviewDate.add(Duration(days: days2NextReviewInt));

    var numRatings = historyOfRatings.length;
    if (numRatings > maxRatings) {
      historyOfRatings = historyOfRatings.sublist(numRatings - maxRatings);
    }
  }
}
