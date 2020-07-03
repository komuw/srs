class SrsException implements Exception {
  final String message;
  const SrsException(this.message);

  @override
  String toString() {
    return "SrsException($message)";
  }
}
