#include <R.h>
#include <Rinternals.h>

extern char* Parse();

SEXP gognparse(SEXP x) {
  return Rf_mkString(Parse(STRING_PTR(x)));
}
