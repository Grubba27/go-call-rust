use clap::Parser;

/// Simple program to greet a person
#[derive(Parser, Debug)]
#[command(author, version, about, long_about = None)]
struct Args {
    /// Name of the person to greet
    #[arg(short, long)]
    name: String,

}

fn main() {
    let args = Args::parse();
    
    println!("Hello, {}!", args.name)
    
}
