import "package:test/test.dart" as tester;
import "package:srs/srs.dart" as srs;

/*
dart format .; pub run test .
 */

void test_card() {
  var c =
      srs.Card("name?", "My name is Kapombe.", srs.Tag("cs", "computer science general knowledge"));

  tester.expect(c.createdAt.isUtc, tester.equals(true));
  tester.expect(c.updatedAt.isUtc, tester.equals(true));
  tester.expect(c.nextReviewDate.isUtc, tester.equals(true));

  tester.expect(c.nextReviewDate.difference(c.createdAt), tester.equals(Duration(days: 1)));

  print("$c");
}

void main() {
  tester.test("test card", test_card, tags: "unit_test");
}
