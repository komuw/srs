import "package:test/test.dart" as tester;
import "package:srs/srs.dart" as srs;

/*
dart format .; pub run test .
 */

void test_srs_algorithm() {
  const test_table_happy = {
    //easies
    [srs.Rating.Easy]: 7.5428796775927704,
    [srs.Rating.Easy, srs.Rating.Easy]: 10.212873916030444,
    [srs.Rating.Easy, srs.Rating.Easy, srs.Rating.Easy]: 14.635400517690599,
    [
      srs.Rating.Easy,
      srs.Rating.Easy,
      srs.Rating.Easy,
      srs.Rating.Easy,
      srs.Rating.Easy,
      srs.Rating.Easy
    ]: 55.03745708646038,

    //mediums
    [srs.Rating.Medium]: 7.112043787939134,
    [srs.Rating.Medium, srs.Rating.Medium]: 8.194714882493026,

    //hards
    [srs.Rating.Hard]: 1.0,
    [srs.Rating.Hard, srs.Rating.Hard]: 1.0,
    [srs.Rating.Hard, srs.Rating.Hard, srs.Rating.Hard, srs.Rating.Hard]: 1.0,

    // easy ending in hard
    [srs.Rating.Easy, srs.Rating.Easy, srs.Rating.Easy, srs.Rating.Hard]: 1.0,
    [
      srs.Rating.Easy,
      srs.Rating.Easy,
      srs.Rating.Easy,
      srs.Rating.Hard,
      srs.Rating.Easy,
    ]: 19.128412543706776,

    //hard ending in easy
    [srs.Rating.Hard, srs.Rating.Hard, srs.Rating.Hard, srs.Rating.Easy]: 6.323243712370701,
    [srs.Rating.Hard, srs.Rating.Hard, srs.Rating.Hard, srs.Rating.Medium]: 6.323243712370701,
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
