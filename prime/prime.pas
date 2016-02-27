{$MODE OBJFPC}
program main;
type IV = class(TObject)
  private 
    res : array of integer;
    used : integer;
  public
    constructor Create; 
    procedure Append(val : integer);
    function At(idx : integer):integer;
    property size : integer read used;
  end;
constructor IV.Create;
  begin
    used := 0;
    SetLength(res,32);
  end;
procedure IV.Append(val : integer);
  var len : integer;
  begin
    len := length(res);
    if used >= len then begin
      if len > 1024 then len := len + len div 8
      else len := len * 2;
      SetLength(res,len);
    end;
    res[used] := val;
    inc(used);
  end;
function IV.At(idx: integer):integer;
  begin
    Result := res[idx];
  end;
var res : IV;
var last : integer = 3;
var v, z : integer;
var prime : boolean;
begin
  res := IV.Create ;
  res.Append(last);
  while(true)do begin
    last := last + 2;
    prime := true;
    for z := 0 to res.size-1 do begin
      v := res.At(z);
  	  if v*v > last then break;
  	  if last mod v = 0 then begin
        prime := false;
        break;
      end
    end;
    if prime then begin
      res.Append(last);
      if res.size mod 100000 = 0 then writeln(last);
      if last > 9999999 then break;
    end
  end
end.