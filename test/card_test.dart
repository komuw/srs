import "package:test/test.dart" as tester;
import "package:srs/srs.dart" as srs;

/*
dart format .; pub run test .
 */

void test_card_creation() {
  var c = srs.Card(
      "name?", "My name is Kapombe.", [srs.Tag("cs", "computer science general knowledge")]);

  tester.expect(c.createdAt.isUtc, tester.equals(true));
  tester.expect(c.updatedAt.isUtc, tester.equals(true));
  tester.expect(c.nextReviewDate.isUtc, tester.equals(true));

  tester.expect(c.nextReviewDate.difference(c.createdAt), tester.equals(Duration(days: 1)));

  var c2 = srs.Card("q?", "a", [
    // tags with same name are treated as same Tag
    srs.Tag("cs", "aa"),
    srs.Tag("algo", "1141"),
    srs.Tag("cs", "xyz"),
    srs.Tag("cs", "3akskq"),
    srs.Tag("cs", "3akskq"),
    srs.Tag("cs", "3akskq"),
    srs.Tag("cs", "3akskq")
  ]);
  // expected is 4 because by default we Card automatically adds `generateDefaultTags`
  tester.expect(c2.tags.length, tester.equals(4));
}

void test_card_tags() {
  var c = srs.Card(
      "name?", "My name is Kapombe.", [srs.Tag("cs", "computer science general knowledge")]);

  List<String> _tags = [];
  var r = c.tags.iterator;
  while (r.moveNext()) {
    _tags.add(r.current.name);
  }
  tester.expect(_tags, tester.equals(["all", "year2020", "cs"]));
}

void test_card_update() {
  var c = srs.Card(
      "Why is it called the dead sea?",
      "Almost nothing lives in it, because its very salty.",
      [srs.Tag("geography", "general knowledge about geography.")]);
  var now = c.createdAt;
  tester.expect(c.nextReviewDate.difference(now), tester.equals(Duration(days: 1)));

  c.update(srs.Rating.Easy);
  tester.expect(c.nextReviewDate.difference(now), tester.equals(Duration(days: 8)));

  c.update(srs.Rating.Easy);
  tester.expect(c.nextReviewDate.difference(now), tester.equals(Duration(days: 18)));

  c.update(srs.Rating.Hard);
  tester.expect(c.nextReviewDate.difference(now), tester.equals(Duration(days: 19)));

  c.update(srs.Rating.Easy);
  tester.expect(c.nextReviewDate.difference(now), tester.equals(Duration(days: 31)));

  c.update(srs.Rating.Easy);
  tester.expect(c.nextReviewDate.difference(now), tester.equals(Duration(days: 50)));

//try and update beyond `Card.maxRatings`
  for (var i = 0; i < srs.Card.maxRatings * 3; i++) {
    c.update(srs.Rating.Easy);
  }
  tester.expect(c.nextReviewDate.difference(now), tester.equals(Duration(days: 97845)));
}

void main() {
  tester.test("test card creation", test_card_creation, tags: "unit_test");
  tester.test("test card tags", test_card_tags, tags: "unit_test");

  tester.test("test card update", test_card_update, tags: "unit_test");
}
