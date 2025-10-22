import { useEffect, useRef } from "react";
import { Terminal } from "xterm";
import "xterm/css/xterm.css";
import {EventsEmit, EventsOn} from "../wailsjs/runtime";

export default function TerminalView() {
  const termRef = useRef<HTMLDivElement>(null);
  const term = useRef<Terminal | null>(null);

  useEffect(() => {
    term.current = new Terminal({
      fontFamily: "monospace",
      theme: { background: "#000000", foreground: "#cccccc" },
      cursorBlink: true,
      scrollback: 1000,
    });
    if (termRef.current && term.current) {
      term.current.open(termRef.current);
      term.current.write("Welcome to TermGo!\r\n");
    }
    term.current.focus();
    term.current.onData((data)=> {
      EventsEmit("pty-input", data);
    });
    const textDecoder = new TextDecoder();
    const unsubscribe = EventsOn("pty-output", (outEvent) => {
      if (typeof outEvent == "string") {
      term.current?.write(outEvent);
      } else {
        term.current?.write((atob(textDecoder.decode(outEvent.data))));
      }
    });
    // We'll connect PTY output here in the next step
    return () => {
      unsubscribe();
      term.current?.dispose();
    };
  }, []);

  return <div ref={termRef} style={{ height: "100%", width: "100%" }} />;
}