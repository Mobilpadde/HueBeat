const int aPin = 0;

void setup(){
  Serial.begin(9600);
}

void loop(){
  int val = analogRead(aPin);

  int mappedVal = map(val, 0, 1023, 1, 100);

  Serial.print(String(mappedVal) + ",");

  delay(500);
}

