with Ada.Containers.Vectors;
with Ada.Integer_Text_IO;
with Ada.Text_IO;
use Ada.Containers;
procedure Prime is
  package IV is new Vectors(Natural,Integer);
  res : IV.Vector;
  Cursor : IV.Cursor;
  last, v : Integer;
  prime : Boolean;
begin
  last := 3;
  IV.Append(res,last);
  Find_Primes:
  loop 
    last := last + 2;
    prime := true;
    Cursor := IV.First(res);
    Check_Divisible:
    while IV.Has_Element(Cursor) loop
      v := IV.Element(Cursor);
      exit Check_Divisible when v*v > last;
      if last mod v = 0 then
        prime := false;
        exit Check_Divisible;
      end if;
      IV.Next(Cursor);
    end loop Check_Divisible;
    if prime then
      IV.Append(res,last); 
      if IV.Length(res) mod 100000 = 0 then
        Ada.Integer_Text_IO.Put(last);
        Ada.Text_IO.New_Line;
      end if;
      exit Find_Primes when last > 9999999;
    end if;
  end loop Find_Primes;
end Prime;

-- -- slower version (16-24 seconds)
--  with Ada.Containers.Vectors;
--  with Ada.Integer_Text_IO;
--  with Ada.Text_IO;
--  use Ada.Containers;
--  procedure Prime is
--    package IV is new Vectors(Natural,Integer);
--    res : IV.Vector;
--    last : Integer;
--    prime : Boolean;
--  begin
--    last := 3;
--    res.Append(last);
--    Find_Primes:
--    loop 
--      last := last + 2;
--      prime := true;
--      Check_Divisible:
--      for v of res loop
--        exit Check_Divisible when v*v > last;
--        if last mod v = 0 then
--          prime := false;
--          exit Check_Divisible;
--        end if;
--      end loop Check_Divisible;
--      if prime then
--        res.Append(last); 
--        if res.Length mod 100000 = 0 then
--          Ada.Integer_Text_IO.Put(last);
--          Ada.Text_IO.New_Line;
--        end if;
--        exit Find_Primes when last > 9999999;
--      end if;
--    end loop Find_Primes;
--  end Prime;
