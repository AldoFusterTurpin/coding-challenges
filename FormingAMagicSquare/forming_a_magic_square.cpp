// ALDO FUSTER TURPIN

#include <bits/stdc++.h>
using namespace std;

typedef vector<int> Row;
typedef vector<Row> Matrix;

//returns the cost (see statement) to go from the matrix "from" to matrix "to"
//we can use "actual_min_cost" to do an early exit.
int costToTransform(const Matrix &from, const Matrix &to, const int &actual_min_cost){
    int cost = 0;

    for (int i = 0; i < from.size(); i++) {
        for (int j = 0; j < from.size(); j++) {
            cost += std::abs(to[i][j] - from[i][j]);

            if (cost > actual_min_cost) {
                return cost;
            }
        }
    }
    return cost;
}


int calculate_min_cost(const Matrix &s, 
                       const vector<Matrix> &latinSquares) {

    //latinSquares[i] is a Matrix                       
    int min_cost = costToTransform(s, latinSquares[0], INT_MAX);

    for (int i = 1; i < latinSquares.size(); i++) {
        min_cost = std::min(min_cost, costToTransform(s, latinSquares[i], min_cost));
    }
    return min_cost;
}


// Complete the formingMagicSquare function below.
int formingMagicSquare(const Matrix &s) {
    vector<Matrix> latin_squares = 
                        //latin_squares[0]
                        {{{8, 1, 6},
                        {3, 5, 7},
                        {4, 9, 2}},

                        //latin_quares[1]
                        {{6, 1, 8},
                        {7, 5, 3},
                        {2, 9, 4}},

                        //latin_quares[2]
                        {{4, 9, 2},
                        {3, 5, 7},
                        {8, 1, 6}},

                        //latin_quares[3]
                        {{2, 9, 4},
                        {7, 5, 3},
                        {6, 1, 8}},

                        //latin_quares[4]
                        {{8, 3, 4},
                        {1, 5, 9},
                        {6, 7, 2}},

                        //latin_quares[5]
                        {{4, 3, 8},
                        {9, 5, 1},
                        {2, 7, 6}},

                        //latin_quares[6]
                        {{6, 7, 2},
                        {1, 5, 9},
                        {8, 3, 4}},

                        //latin_quares[7]
                        {{2, 7, 6},
                        {9, 5, 1},
                        {4, 3, 8}}};

    return calculate_min_cost(s, latin_squares);
}

int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    vector<vector<int>> s(3);
    for (int i = 0; i < 3; i++) {
        s[i].resize(3);

        for (int j = 0; j < 3; j++) {
            cin >> s[i][j];
        }

        cin.ignore(numeric_limits<streamsize>::max(), '\n');
    }

    int result = formingMagicSquare(s);

    fout << result << "\n";

    fout.close();

    return 0;
}
