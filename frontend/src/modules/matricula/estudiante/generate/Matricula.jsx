import { jsPDF } from "jspdf";

export default function Matricula() {
  const generarPDF = () => {
    const doc = new jsPDF();
    doc.text("a", 10, 10);
    doc.save("matricula.pdf");
  };

  return (
    <button
      onClick={generarPDF}
      className="ml-6 bg-gray-200 text-black text-xs font-semibold rounded-md px-4 py-1 hover:bg-gray-300"
    >
      Generar Matrícula
    </button>
  );
}
