class Solution {
   public static int solution(int n) {
        int answer = 0;
        int start = 1;
        int end = 1;
        int sum = 1;

        while (start <= n) {
            if (sum < n) {
                end += 1;
                sum += end;
            } else if ( sum == n ){
                answer += 1;
                start += 1;
                sum -= start - 1;
            } else {
                start += 1;
                sum -= start - 1;
            }
        }

        return answer;
    }
}