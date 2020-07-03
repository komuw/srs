import "dart:math" as math;

class Tag {
  String name;
  String description;

  // all the datetimes should be in UTC
  late DateTime createdAt;
  late DateTime updatedAt;

  Tag(this.name, this.description) {
    var now = DateTime.now().toUtc();
    createdAt = now;
    updatedAt = now;
  }

  @override
  String toString() {
    var _nSubst = name.substring(0, math.min(10, name.length));
    var _dSubst = description.substring(0, math.min(10, description.length));

    return "Tag(N:$_nSubst, D:$_dSubst)";
  }
}

/// The system will have some default tags.
/// One such is the `all` tag that will be added to all cards.
/// Another is a tag for each year that this system has been in existence.
/// Each card will get a tag added for the year it was created.
Set<Tag> generateDefaultTags() {
  var thisYear = DateTime.now().toUtc().year;
  var defaultTags = [
    Tag("all", "tag availed/added to all cards."),
    Tag("year$thisYear", "tag for year $thisYear")
  ];

  return Set.from(defaultTags);
}
