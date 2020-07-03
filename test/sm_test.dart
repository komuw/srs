import "package:test/test.dart" as tester;
import "package:srs/srs.dart" as srs;

/*
dart format .; pub run test .
 */

void test_srs_algorithm() {
  const test_table_happy = {
    [srs.Rating.Easiest, srs.Rating.Easiest, srs.Rating.Easiest]: 17.226936248077187,
    [
      srs.Rating.Easiest,
      srs.Rating.Easiest,
      srs.Rating.Easiest,
      srs.Rating.Easiest,
      srs.Rating.Easiest,
      srs.Rating.Easiest
    ]: 84.91822762093749,
    [srs.Rating.Hardest, srs.Rating.Hard, srs.Rating.Medium, srs.Rating.Hardest]: 1.0,
    [srs.Rating.Hardest]: 1.0,
    [srs.Rating.Hard]: 1.0,
    [
      srs.Rating.Easy,
      srs.Rating.Easy,
      srs.Rating.Easy,
      srs.Rating.Easy,
      srs.Rating.Easy,
      srs.Rating.Easy,
    ]: 55.03745708646038,
    [
      srs.Rating.Easy,
      srs.Rating.Easy,
      srs.Rating.Easy,
      srs.Rating.Easy,
      srs.Rating.Easy,
      srs.Rating.Easy,
      srs.Rating.Hard
    ]: 1.0,
    [srs.Rating.Medium, srs.Rating.Medium, srs.Rating.Medium, srs.Rating.Medium]: 15.895881282351272
  };

  test_table_happy.forEach((key, value) {
    tester.expect(srs.sm2(key), tester.equals(value));
  });

//  const test_table_sad = {
//    [-1]: AssertionError
//  };
//  test_table_sad.forEach((key, value) {
//    tester.expect(srs.sm2(key), tester.throwsA(AssertionError("ll")));
//  });
}

void main() {
  tester.test("test srs algorithm", test_srs_algorithm, tags: "unit_test");
}
