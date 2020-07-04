import "package:test/test.dart" as tester;
import "package:srs/srs.dart" as srs;

/*
dart format .; pub run test .
 */

void test_card_creation() {
  var c =
      srs.Card("name?", "My name is Kapombe.", srs.Tag("cs", "computer science general knowledge"));

  tester.expect(c.createdAt.isUtc, tester.equals(true));
  tester.expect(c.updatedAt.isUtc, tester.equals(true));
  tester.expect(c.nextReviewDate.isUtc, tester.equals(true));

  tester.expect(c.nextReviewDate.difference(c.createdAt), tester.equals(Duration(days: 1)));
}

void test_card_tags() {
  var c =
      srs.Card("name?", "My name is Kapombe.", srs.Tag("cs", "computer science general knowledge"));

  List<String> _tags = [];
  var r = c.tags.iterator;
  while (r.moveNext()) {
    _tags.add(r.current.name);
  }
  tester.expect(_tags, tester.equals(["all", "year2020", "cs"]));
}

void main() {
  tester.test("test card creation", test_card_creation, tags: "unit_test");
  tester.test("test card tags", test_card_tags, tags: "unit_test");
}
