import "dart:math" as math;
import "./exceptions.dart" as e;

class Tag {
  final String name;

  // all the datetimes should be in UTC
  late final DateTime createdAt;
  late final DateTime updatedAt;

  Tag(this.name) {
    if (name.contains(" ")) {
      throw e.SrsException("tag's name should be one word");
    }
    var now = DateTime.now().toUtc();
    createdAt = now;
    updatedAt = now;
  }

  @override
  String toString() {
    var _nSubst = name.substring(0, math.min(15, name.length));

    return "Tag(N:$_nSubst)";
  }

  @override
  bool operator ==(other) {
    return other is Tag && other.name == name;
  }

  @override
  int get hashCode {
    return name.hashCode;
  }
}

/// The system will have some default tags.
/// One such is the `all` tag that will be added to all cards.
/// Another is a tag for each year that this system has been in existence.
/// Each card will get a tag added for the year it was created.
Set<Tag> generateDefaultTags() {
  var thisYear = DateTime.now().toUtc().year;
  var defaultTags = [Tag("all"), Tag("year$thisYear")];

  return Set.from(defaultTags);
}
