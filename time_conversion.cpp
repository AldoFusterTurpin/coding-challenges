// ALDO FUSTER TURPIN

#include <bits/stdc++.h>

using namespace std;

// offset can be AM or PM
string ConvertHoursTo24Format(const string &hour, const string &offset) {
  if (offset == "AM") {
    return hour == "12" ? "00" : string(hour);
  }

  // else (offset == "PM")
  int int_hour = stoi(hour) + 12;
  return hour == "12" ? "12" : to_string(int_hour);
}
/*
 * 01 <= hh <= 12
 * 00 <= mm, ss <= 59
 */
string timeConversion(string s) {

  string hour = "";
  hour += s[0];
  hour += s[1];

  string offset = "";
  offset += s[8];
  offset += s[9];

  return ConvertHoursTo24Format(hour, offset) + s[2] + s[3] + s[4] + s[5] + s[6] + s[7];
}

int main() {
  ofstream fout(getenv("OUTPUT_PATH"));

  string s;
  getline(cin, s);

  string result = timeConversion(s);

  fout << result << "\n";

  fout.close();

  return 0;
}
