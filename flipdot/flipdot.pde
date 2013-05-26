import processing.net.*;

int x, y;
int scale = 8;
int i, j;
boolean[][] flipped;
Client c;
boolean change;

int h = 16, w = 112;


void setup()
{
  size(112*scale, 16*scale);
  flipped = new boolean[w][h];
  c = new Client(this, "192.168.23.222", 6210);
  c.write("c"); 
}

void draw() {
  background(0);
  noStroke();
  
  x = (mouseX - (mouseX % scale)) / scale;
  y = (mouseY - (mouseY % scale))/ scale;
  
  println("x, y: " + x + " " + y);
  
  if( mousePressed &&
     (x < w && x >= 0) &&
     (y < h && y >= 0) ) {
    flipped[x][y] = mouseButton == LEFT;
  
  
    if ( flipped[x][y] ) {
      c.write("1" + char(x) + char(y));
    } else {
      c.write("0" + char(x) + char(y));
    }
  }
  
  for(i = 0; i < w; i++) {
    for(j = 0; j < h; j++) {
      if( flipped[i][j] ) {
        fill(204, 255, 0);
      } else {
        fill(0, 0, 0);
      }
      
      rect(i*scale, j*scale, scale, scale);
    }
  }
  
  fill(92, 92, 92);
  rect(x*scale, y*scale, scale, scale);
}

