const int aPin = 0;

void setup(){
  Serial.begin(9600);
}

void loop(){
  int val = analogRead(aPin);

  Serial.print(String(val) + ",");

  delay(500);
}

