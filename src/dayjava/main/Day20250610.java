package main;

public class Day20250610 {

    public static void main(String[] args) {

    }

    public int maxDifference(String s) {
        int[] cnt = new int[26];

        for (byte b : s.getBytes()) {
            cnt[b - 'a']++;
        }

        int a1 = 0;
        int a2 = 0;

        for (int count : cnt) {
            if (count == 0) {
                continue;
            }
            if (count % 2 == 1 && count > a1) {
                a1 = count;
                continue;
            }

            if (count % 2 == 0 && (a2 == 0 || count < a2)) {
                a2 = count;
            }

        }
        return a1 - a2;
    }
}
