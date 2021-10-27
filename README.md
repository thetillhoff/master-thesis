# `Master thesis on extending Infrastructure-as-Code to bare-metal`

Infrastructure as Code (IaC) is a trend that applies software development techniques to infrastructure.
Most currently available tools in that area require a (cloud) backend that is available at all times and listens on API-calls like "I want X servers to be provisioned as Y".
But these backends have to be provisioned as well.
And still, at some level, hardware needs to be provisioned.
A simple "Just give me a server X with Y" does not work here - or does it?
This thesis investigates whether bare-metal machines can be provisioned on-demand with orchestration software that also runs on-demand.
By doing so, the initially required manual setup could be eliminated or at least dramatically reduced.
To describe the virtual and physical infrastructure with a consistent language, a domain-specific language is selected and a systematic approach for integrating the description of hardware is developed.

<!-- ultimate goal: Publish as conference paper -->

<!-- ## Exposé ("Arbeitsplan")

## Vorläufiger Titel
Leveraging (immutable) (open-source) infrastructure-as-code for enterprise environments
infrastructure lifecycle management
infrastructure configuration management
infrastructure management system
comparison of infrastructure-as-code dsls and tools (and finding/creating a compromise)
designing a pluggable tosca orchestrator & simplify user-onboarding -->

<!-- ## Relevanz des Themas (aus praxisorientierter und wissenschaftlicher Perspektive)
Praxis: Was ist das Ziel, was gibt es dafür und was ist am Ansatz in dieser Thesis so besonders?
Wissenschaftlich: Relevanz von Cluster-API, OpenStack, tinkerbell, vSphere

## Zusammenfassung des Stands der aktuellen Forschung zum Themenbereich, wichtige Begriffe & Forschungslücken -->


<!-- ## Forschungsfragen und Ziele
Konkretisierung anhand von:
- Anwendungsbereich "untersucht am Beispiel von"
- Nischenaspekte "vor dem Hintergrund von"
- Blickwinkel "auf der Mitarbeiterebene / auf der technischen Ebene"
- Beziehungen -->

<!-- - Goal is to allow dynamic provision of k8s clusters for users - including storage
  - why multitenancy via multiple clusters
- From the base up; not cloud-vendor specific!
- Are VMs dead (are they even needed / the case for bare-metal)
  - comparison of bare-metal approach vs vSphere and/or OpenStack approach
  - constraints like
    - Workload comparison; are there workloads which cannot run in containers and require VMs?
    - minimum machine size defines minimum cluster size and therfore introduces unused resources
  - -> VMs make sense! What about their overhead? They need "zone/node affinity" as well
- OpenStack does a lot of those things, vSphere as well
  - is there a more lightweight or open source approach? A lot features aren't needed -->


<!-- - one mgmt server per rack for ipmi control, os distribution -> can this be done with top-of-rack-switch?
- instead of special network topologies, scheduling stuff, just let all bm-hosts communicate with all other bm-hosts
  - consider zone-affinity when allocating
  - try to allocate "closest fitting" node
- differentiate between bm-cluster and vm-cluster
  - in bm-cluster everyone can talk with everyone, zones can be nested, zones can be distinguished via address-spaces
  - in vm-cluster vlans & routes must be configured, vms should be placed close together (still consider node/zone affinity) to reduce hops
  - chicken and egg bootstrapping problem (security-wise, control-cluster-wise)
- two options for automated on-prem/non-cloud datacenter:
  - openstack
  - layers
    - hypervisor provisioning (on-demand with IPMI)
    - vm provisioning (on-demand via hypervisor API, f.e. vSphere)
    - clustering via k8s, or, if necessary, allow direct access to vms for some applications (f.e. Active Directory)
- outlook: application deployment; helm, kustomize, ...
- how to automate DNS
- stability of openstack? why so many problems at CaaS? -->


<!-- - there was an example where cloud platform migration was necessary for an application, and the lead developed "magic" - something similar
  - https://dzone.com/articles/reducing-your-cloud-costs-by-90
  - https://dzone.com/articles/nomadic-cloud-systems
  - https://polterguy.github.io/
- multi-cloud where everything-is-a-service and is based on a below layer -> lower layers can be translated between different cloud-providers -> making the whole thing cloud-agnostic
- when a cloud has a super cool feature and the others don't, but you want to use it, migrate. If others introduce it later, and are cheaper, migrate.
- the whole cloud-, container-, k8s-, iac-collection of things is about migration between environments -> the goal should be to make migration as easy as possible. As a cloud provider, I want to lock the user in, and catch foreigners as well, but make it subtle enough they are not feared away or give them advantages like superb prices.
  - migration between dev, staging and prod, migration between dev-machine and dev-machine as well as dev-machine and server, migration to a larger scale system (by scaling lol).
  - I adopt some new tech, f.e. docker (-> by writing docker files), now I want to migrate that somewhere else, so please make that easy goddamnit!
- migration-path ==> execution path -->

<!-- - with an everything-as-a-service-architecture, this should be easy!
  - have one language-fits-them-all, f.e. TOSCA/whatever
  - create an everything-as-a-service layered architecture
    - have a translation service between different providers, or be able to translate top level language into all different provider-dsls
  - ? use terraform as underlying base
- "micrate" (migrate, cloud, crate?)
- all "cloud-agnostic"-tools are just capable of running a scenario for different providers, but migration is still not possible, because of the provider-specific stuff. So providers seem to implement the same stuff differently making migrations unnecessary difficult.
- "truly cloud-agnostic iac-tooling" by implementing everything-as-a-service -->

<!-- scaling smaller is better: https://dzone.com/articles/cost-optimization-strategies-for-compute-instances
cost optimization tools: https://dzone.com/articles/zoom-spotify-and-others-slashed-their-cloud-costs
cloud vs serverless: https://dzone.com/articles/serverless-computing-vs-cloud-computing -->

<!-- - read chapter 10ff of simple-profile
- read chapter 13ff of simple-profile
- don't reinvent the wheel, allow other standards as artifacts, like openapi as api-spec -->

<!-- ## Vision
The goal: [`live-demo-vision.md`](./live-demo-vision.md) -->

<!-- ## Zeitplan
- read papers, copy notes for `related work` and `background` (7-22p -> 4w)
- background (7-22p -> 2w a 7d; ~1p/d)
- introduction (3-4p -> 1w)
- design & implementation (5-10p -> 6w)
- evaluation (5-10p -> 2w)
- discussion (5-10p -> 2w)
- conclusion (2-4p -> 1w)
- abstract (1-2p -> 1w)
-> https://github.com/thetillhoff/master-thesis/milestones?direction=asc&sort=due_date&state=open -->

<!-- ## Literature sources
- google scholar
- ieee.org/dl
- acm.org/dl
- gi.de/service/digitale-bibliothek -->

<!-- ## Notes for example toc in `notes.yml`
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
``` -->

<!-- ## finishing touches
- [x] Check comments in this readme for additional tasks (english and german)
- [x] for checking text, make sure to stick to either british or american english
- [ ] Check text with hemingwayapp https://hemingwayapp.com/
- [ ] Check text with languagetool https://languagetool.org/
- [ ] Check text with scribens https://www.scribens.com/
- [ ] Check the text with reverso https://www.reverso.net/spell-checker/english-spelling-grammar/
- [x] Check line-endings
- [ ] Resolve latex-warnings on f.e. overfull lines etc.
- [x] Use italic to emphasise, NEVER bold, capitals or underlines as those break the flow.
- [x] Ensure unified use of pagebreaks before headlines (level 1, 2 etc)
- [x] The Elements of Style: http://www.bartleby.com/141/
- [x] Remove the `\include{content/demo}` from main
- [x] Check for highlighted contents
- [ ] Search for `%TODO`, `% TODO` and simply `TODO`
- [x] Check for sidenodes
- [x] Search for '[' and ']', check formatting of code listings
- [x] Check the date on the titlepage -->

<!-- ## Häufige Fehler
- Es fehlt der "rote Faden": Die Arbeit lässt keine Logik in Bezug auf die Fragestellung erkennen. Der Zusammenhang zwischen den Kapiteln wird nicht deutlich.
- Der Leser sollte zu jedem Zeitpunkt verstehen "warum" Sie etwas tun. Versuchen Sie Ihre Überlegungen klar darzulegen. Wenn Sie z.B. eine bestimmte Methodik verwenden, erklären Sie, warum Sie sich für diese Methodik und nicht für eine andere entschieden haben. Wenn Sie z.B. ein Bewertungsschema verwenden und für die Erfüllung von Kriterien Punkte vergeben, erklären Sie, wie Sie diese Punkte ermittelt haben und warum Sie diese auf genau diese Art ermittelt haben. Sobald sich im Kopf des Lesers ein "Häh? -Warum jetzt das?" bildet, haben Sie versäumt, wichtige Informationen zu vermitteln.
- Mangelnde Reflexion des Stoffes. Sie sollen eine Fragestellung / ein Problem lösen! Ordnen Sie Ihren Stoff ein und beurteilen Sie ihn wissenschaftlich. Also nicht nur Literaturquellen zusammensuchen und wiedergeben sondern eigenen Erkenntnisgewinn erzeugen.
- Bitte achten Sie auf Lesefreundlichkeit! Lange, verschachtelte Sätze mit einer Vielzahl an Fremdwörtern sind kein Kriterium für wissenschaftliche Kompetenz. Kurze, klare Sätze sind für Sie einfacher zu schreiben und für den Leser erheblich leichter zu verstehen....und was auch immer wieder passiert:
- Eine Untergliederung mit nur einem Gliederungspunkt. Dies mag zwar manchmal praktisch erscheinen, ist aber falsch!
- Falsche Zeilenumbrüche – unbedingt kontrollieren... und ganz wichtig...
- Achten Sie bitte auf Rechtschreibung und Grammatik – auch das fließt in die Note ein. Immer wieder werden Arbeiten vorgelegt, bei denen ich 10 Fehler pro Seite anstreiche. Dies vermittelt den Eindruck einer schlampig angefertigten Arbeit. Wenn Sie Probleme mit deutscher Rechtschreibung und Grammatik haben, suchen Sie sich einen Freund/Freundin/Familienmitglied für die Korrektur. -->

<!-- ## ...und noch ein paar Tipps für den Start
Es ist hilfreich, sich zunächst intensiv mit der Einleitung auseinanderzusetzen. Manchmal liest man in Ratgebern "schreib die Einleitung zuletzt". Das mag für die finale Fassung Ihrer Arbeit sinnvoll sein, inhaltlich müssen Sie sich aber unbedingt zu Beginn damit auseinandersetzen. Was genau ist denn das Ziel der Arbeit? Was soll das Ergebnis der Arbeit sein? Was soll Ihre Arbeit bewirken? Welches Problem möchten Sie lösen? Und wie können Siedieses Ergebnis am sinnvollsten erreichen? Und warum ist die ganze Arbeit überhaupt wichtig? Welche Relevanz hat das Thema? Zu Beginn ist es oft schwierig einen Einstieg in die Literaturrecherche zu finden. Es bietet sich an, mit Lehrbüchern zu den für Sie relevanten Themengebieten zu beginnen und über die Literaturverzeichnisse weiter zu suchen. Und bevor Sie mal nicht weiter wissen oder unsicher über das weitere Vorgehen sind, kontaktieren Sie Ihren betreuenden Professor / Ihre betreuende Professorin –dafür sind wir da. -->
