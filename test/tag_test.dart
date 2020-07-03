import "package:test/test.dart" as tester;
import "package:srs/srs.dart" as srs;

/*
dart format .; pub run test .
 */

void test_tag() {
  var c = srs.Tag("all", "tag availed/added to all cards.");

  tester.expect(c.createdAt.isUtc, tester.equals(true));
  tester.expect(c.updatedAt.isUtc, tester.equals(true));

  var defaultTags = srs.generateDefaultTags();

  print("defaultTags:: $defaultTags");
}

void main() {
  tester.test("test tag", test_tag, tags: "unit_test");
}
