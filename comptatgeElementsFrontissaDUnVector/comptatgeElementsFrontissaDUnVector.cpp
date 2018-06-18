//ALDO FUSTER TURPIN
 #include <iostream>
 #include <vector>
using namespace std;


/* Pre: v vector of int's */
/* Post: the return vector is "leftSum" -> leftSum[i] is the left sum of v[i]
 left sum means the sum of all elements from 0 to i-1.
 ie: leftSum[i] = sum of v[0 ... i-1] */
vector<int> calculateLeftSum(const vector<int> &v)
{
    int n = v.size();
    if (n > 0) {
        //leftSum[i] is the left sum of the element at v[i]
        vector<int> leftSum(n);

        leftSum[0] = 0;

        for (int i = 1; i < n; i++) {
            leftSum[i] = v[i-1] + leftSum[i-1];
        }

        /************ Debug ***********/
        // cout << "leftSum: ";
        // for (int i = 0; i < n; i++) cout << leftSum[i] << " ";
        // cout << endl;
        /****************************/
        return leftSum;
    }

    //if v has no elements we return an empty vector (better than NULL)
    else return vector<int>(0);
}


int calculateFrontisses(const vector<int> &v, const vector<int> & leftSum)
{
    //number of "elements frontissa"
    int counter = 0;
    int n = v.size();

    if (n < 1) return 0;

    //rightSum[i] is the right sum of the element at v[i]
    vector<int> rightSum(n);

    rightSum[n-1] = 0;

    for (int i = n-2; i >= 0; i--) {
        rightSum[i] = v[i+1] + rightSum[i+1];
        if (rightSum[i] - leftSum[i] == v[i]){
            ++counter;
        }
    }

    if (rightSum[n-1] - leftSum[n-1] == v[n-1]) {
        ++counter;
    }


    /********* Debug *********/
    // cout << "rightSum: ";
    // for (int i = 0; i < n; i++) cout << rightSum[i] << " ";
    /****************************/

    return counter;
}


/* Pre: cert */
/* Post: el resultat es el nombre d'elements frontissa de v */
int comptatge_frontisses(const vector<int> &v)
{
    if (v.size() < 1) return 0;

    vector<int> leftSum = calculateLeftSum(v);
    return calculateFrontisses(v, leftSum);
}




int main()
{
    int n;
    cin >> n;

    vector<int> v(n);

    for (int i = 0; i < n; i++) cin >> v[i];

    int result = comptatge_frontisses(v);
    cout << endl << "nelements frontissa: " << result << endl;

    return 0;
}
