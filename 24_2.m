syms a b c d e f x y z

d = solve([359781776524153-44*x==a+d*x, 312705660279075-125*x==b+e*x, 236728636905923+18*x==c+f*x, 276481733510955+35*y==a+d*y, 270867065789660+20*y==b+e*y, 273768862611813+33*y==c+f*y, 189537654420103+102*z==a+d*z, 292422605212995-15*z==b+e*z, 333617095281945-14*z==c+f*z], [a b c d e f x y z]);
disp(d.a+d.b+d.c)
