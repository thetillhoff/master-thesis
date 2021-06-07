# `master thesis`

**University/Dep.**: Ulm University of Applied Sciences, faculty of computer science

## Exposé ("Arbeitsplan")

### Vorläufiger Titel
Leveraging (immutable) (open-source) infrastructure-as-code for enterprise environments
infrastructure lifecycle management
infrastructure configuration management
infrastructure management system

### Relevanz des Themas (aus praxisorientierter und wissenschaftlicher Perspektive)
Praxis: Was ist das Ziel, was gibt es dafür und was ist am Ansatz in dieser Thesis so besonders?
Wissenschaftlich: Relevanz von Cluster-API, OpenStack, tinkerbell, vSphere

### Zusammenfassung des Stands der aktuellen Forschung zum Themenbereich, wichtige Begriffe & Forschungslücken


### Forschungsfragen und Ziele
Konkretisierung anhand von:
- Anwendungsbereich "untersucht am Beispiel von"
- Nischenaspekte "vor dem Hintergrund von"
- Blickwinkel "auf der Mitarbeiterebene / auf der technischen Ebene"
- Beziehungen

### Methodisches Vorgehen / Implementierung
The goal: [`live-demo-vision.md`](./live-demo-vision.md)

### Zeitplan
- read papers, copy notes for `related work` and `background` (7-22p -> 4w)
- background (7-22p -> 2w a 7d; ~1p/d)
- introduction (3-4p -> 1w)
- design & implementation (5-10p -> 6w)
- evaluation (5-10p -> 2w)
- discussion (5-10p -> 2w)
- conclusion (2-4p -> 1w)
- abstract (1-2p -> 1w)
-> https://github.com/thetillhoff/master-thesis/milestones?direction=asc&sort=due_date&state=open

### Literaturrecherche
- google scholar
- ieee.org/dl
- acm.org/dl
- gi.de/service/digitale-bibliothek

## Notes for example toc in `notes.yml`
```yaml
4 pillars:
 - Hardware(-instances)
 - provisioning
 - config mgmt
 - Security & ID/User mgmt
 - RBAC
 - zero trust
 - multi-tenant environments
 - Storage / Persistance
 - Rook
 - Networking
 - (Service-)Mesh
subtopics:
 - BYOD, COD, UOD
 - bootstrapping 0-100
 - 2nd day
 - OSS, support/SLA/stability/maturity
```

## finishing touches
- for checking text, make sure to stick to either british or american english
- Check text with hemingwayapp https://hemingwayapp.com/
- Check text with languagetool https://languagetool.org/
- Check text with scribens https://www.scribens.com/
- Check line-endings
- Check use of paragraphs instead of empty lines
- Use italic to emphasise, NEVER bold, capitals or underlines as those break the flow.
- Ensure unified use of pagebreaks before headlines (level 1, 2 etc)
- Remove larger unused/commented blocks
- The Elements of Style: http://www.bartleby.com/141/
- Remove the `\include{content/demo}` from main
- Check for highlighted contents
- Search for `%TODO`
- Check for sidenodes
- Check the date on the titlepage

### Häufige Fehler
- Es fehlt der „rote Faden“: Die Arbeit lässt keine Logik in Bezug auf die Fragestellung erkennen. Der Zusammenhang zwischen den Kapiteln wird nicht deutlich.
- Der Leser sollte zu jedem Zeitpunkt verstehen „warum“ Sie etwas tun. Versuchen Sie Ihre Überlegungen klar darzulegen. Wenn Sie z.B. eine bestimmte Methodik verwenden, erklären Sie, warum Sie sich für diese Methodik und nicht für eine andere entschieden haben. Wenn Sie z.B. ein Bewertungsschema verwenden und für die Erfüllung von Kriterien Punkte vergeben, erklären Sie, wie Sie diese Punkte ermittelt haben und warum Sie diese auf genau diese Art ermittelt haben. Sobald sichim Kopf des Lesers ein „Häh? -Warum jetzt das?“ bildet, haben Sie versäumt, wichtige Informationen zu vermitteln.
- Mangelnde Reflexion des Stoffes. Sie sollen eine Fragestellung / ein Problem lösen! Ordnen Sie Ihren Stoff ein und beurteilen Sie ihn wissenschaftlich. Also nicht nur Literaturquellen zusammensuchen und wiedergeben sondern eigenen Erkenntnisgewinn erzeugen.
- Bitte achten Sie auf Lesefreundlichkeit!Lange, verschachtelte Sätze mit einer Vielzahl an Fremdwörtern sind kein Kriterium für wissenschaftliche Kompetenz. Kurze, klare Sätze sind für Sie einfacher zu schreiben und für den Leser erheblich leichter zu verstehen....und was auch immer wieder passiert:
- Eine Untergliederung mit nur einem Gliederungspunkt. Dies mag zwar manchmal praktisch erscheinen, ist aber falsch!
- Falsche Zeilenumbrüche –unbedingt kontrollieren...undganz wichtig...
- Achten Sie bitte auf Rechtschreibung und Grammatik–auch das fließt in die Note ein. Immer wieder werden Arbeiten vorgelegt, bei denen ich 10 Fehler pro Seite anstreiche. Dies vermittelt den Eindruck einer schlampig angefertigten Arbeit. Wenn Sie Probleme mit deutscher Rechtschreibung und Grammatik haben, suchen Sie sich einen Freund/Freundin/Familienmitglied für die Korrektur.

### ...und noch ein paar Tipps für den Start
Es ist hilfreich, sich zunächst intensiv mit der Einleitung auseinanderzusetzen. Manchmal liest man in Ratgebern „schreib die Einleitung zuletzt“. Das mag für die finale Fassung Ihrer Arbeit sinnvoll sein, inhaltlich müssen Sie sich aber unbedingt zu Beginn damit auseinandersetzen. Was genau ist denn das Ziel der Arbeit? Was soll das Ergebnis der Arbeit sein? Was soll Ihre Arbeit bewirken? Welches Problem möchten Sie lösen? Und wie können Siedieses Ergebnis am sinnvollsten erreichen? Und warum ist die ganze Arbeit überhaupt wichtig? Welche Relevanz hat das Thema? Zu Beginn ist es oft schwierig einen Einstieg in die Literaturrecherche zu finden. Es bietet sich an, mit Lehrbüchern zu den für Sie relevanten Themengebieten zu beginnen und über die Literaturverzeichnisse weiter zu suchen. Und bevor Sie mal nicht weiter wissen oder unsicher über das weitere Vorgehen sind, kontaktieren Sie Ihren betreuenden Professor / Ihre betreuende Professorin –dafür sind wir da.
