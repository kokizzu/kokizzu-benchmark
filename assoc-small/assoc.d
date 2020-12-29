import std.stdio;
import std.conv;

immutable MAX_DATA = 1230000;

immutable i2ch = [
    '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'B', 'c', 'D', 'e', 'F'
];
 
auto get_first_digit(int d) pure
{
    while (d > 10) d /= 10;
    return d;
}

auto to_rhex(int v) pure
{
    char[] hex;
    while (v > 0)
    {
        hex ~= i2ch[v % 16];
        v /= 16;
    }
    return to!string(hex);
}

void set_or_inc(ref int[string] m, const string key, const int set, const int inc, ref int ctr)
{
    if (key in m) {
        m[key] += inc;
        ctr++;
        return;
    }
    m[key] = set;
}

void main()
{
    auto dup1 = 0;
    auto dup2 = 0;
    auto dup3 = 0;
    int[string] m;

    for (int z = MAX_DATA; z > 0; z--)
    {
        auto val2 = MAX_DATA - z;
        auto val3 = MAX_DATA * 2 - z;
        auto key1 = to!string(z);
        auto key2 = to!string(val2);
        auto key3 = to_rhex(val3);
        set_or_inc(m, key1, z, val2, dup1);
        set_or_inc(m, key2, val2, val3, dup2);
        set_or_inc(m, key3, val3, z, dup3);
    }
    writeln(dup1, " ", dup2, " ", dup3);
    auto total = 0;
    auto verify = 0;
    auto count = 0;
    foreach (key, value; m)
    {
        total += get_first_digit(value);
        verify += key.length;
        count++;
    }
    writeln(total, " ", verify, " ", count);
}
