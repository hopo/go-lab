package a17_InnerClasses;

public class Outer {
	static class Nested_Demo {
		public void my_method() {
			System.out.println("This si my nested class");
		}
	}
	
	public static void main(String[] args) {
		Outer.Nested_Demo nested = new Outer.Nested_Demo();
		nested.my_method();
	}
}
