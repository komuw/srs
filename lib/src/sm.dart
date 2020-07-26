import "dart:math" as math;

// The original sm2 algo has 6 ratings. I have wittled them down to 3.
// https://gist.github.com/doctorpangloss/13ab29abd087dc1927475e560f876797
enum Rating { Hard, Medium, Easy }

// TODO: look at https://www.idorecall.com/
// They allow you to upload your reading materials(books, pdfs, webpages, videos, youtube etc)
// and then u can link those into your cards.
// When reviewing cards, u can click to go to the exact source, at the exact page(or time if it is a youtube video)

num sm2(List<Rating> historyOfRatings) {
  /*
          Returns the number of days to delay the next review of an item by, fractionally, based on the history of answers to
          a given question, where
          x == 0: Incorrect, Hardest
          x == 1: Incorrect, Hard
          x == 2: Incorrect, Medium
          x == 3: Correct, Medium
          x == 4: Correct, Easy
          x == 5: Correct, Easiest

          @param x The history of answers in the above scoring.
          @param theta When larger, the delays for correct answers will increase.

          https://gist.github.com/doctorpangloss/13ab29abd087dc1927475e560f876797
  */
  const double a = 6.0;
  const double b = -0.8;
  const double c = 0.28;
  const double d = 0.02;
  const double theta = 0.2;

  // TODO: if item>2 set it to 2. if less than 0, set it to 0
  historyOfRatings.forEach((item) {
    if (item.index > 2) {
      throw AssertionError("item; `$item` should not be greater than 2");
    }
    if (item.index < 0) {
      throw AssertionError("item; `$item` should not be less than 0");
    }
  });

  List<bool> correct_x = [false];
  historyOfRatings.forEach((item) {
    if (item.index >= 1) {
      correct_x.add(true);
    } else {
      correct_x.add(false);
    }
  });
  correct_x.removeAt(0);

  // If you got the last question incorrect, just return 1
  if (correct_x[correct_x.length - 1] == false) {
    return 1.0;
  }

  // Calculate the latest consecutive answer streak
  int num_consecutively_correct = 0;
  correct_x.reversed.forEach((i) {
    if (i == true) {
      num_consecutively_correct = num_consecutively_correct + 1;
    }
  });

  List<double> _temp = [0.00];
  historyOfRatings.forEach((i) {
    // since original algorithm had 6 states; this was originally `_temp.add(b + c * i.index + d * i.index * i.index)`
    // however, we have to extrapolate our 3 states to look like 6. so we add a multiplier.
    // the original highest state was 5(zero indexed), our highest is 2. so multiplier is 5/2
    var _multiplier = i.index * 2.00;
    _temp.add(b + c * _multiplier + d * _multiplier * _multiplier);
  });
  _temp.removeAt(0);

  var inner_sum = 0.0;
  _temp.forEach((i) {
    inner_sum += i;
  });

  var _max = math.max(1.3, 2.5 + inner_sum);

  return a * math.pow(_max, theta * num_consecutively_correct);
}

void main() {
  var x = [
    Rating.Easy,
    Rating.Easy,
    Rating.Easy,
  ];
  num res = sm2(x);
  print("res: ${res}");
}
