import { useEffect, useRef } from "react";
import { Terminal } from "xterm";
import "xterm/css/xterm.css";

export default function TerminalView() {
  const termRef = useRef(null);

  useEffect(() => {
    const term = new Terminal({
      fontFamily: "monospace",
      theme: { background: "#000000", foreground: "#cccccc" },
      cursorBlink: true,
      scrollback: 1000,
    });

    term.open(termRef.current);
    term.write("Welcome to TermGo!\r\n");

    // We'll connect PTY output here in the next step
    return () => term.dispose();
  }, []);

  return <div ref={termRef} style={{ height: "100%", width: "100%" }} />;
}
