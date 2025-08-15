# Neural Quest Interface (NQI)

## Transform Your ADHD Brain's Endless Ideas Into Finished Projects

Neural Quest Interface is a revolutionary project management system designed
specifically for neurodivergent minds. Built on Quest-Based Agile principles,
NQI provides external scaffolding for executive functions, helping you bridge
the gap between knowing what to do and actually doing it.

**🎯 Core Philosophy**: Work _with_ your ADHD brain, not against it.

---

## ✨ What Makes NQI Different

### **Six-Column Visual Workflow**

- **🌱 Idea Greenhouse**: Capture ideas without pressure to act immediately
- **📚 Quest Log**: Your backlog of potential projects and tasks
- **🎯 This Cycle's Quest**: Single Epic focus for 1-2 weeks
- **📋 Next Up**: Ready-to-start tasks (maximum 5)
- **⚡ In Progress**: Strict WIP limit of 1 for sustained focus
- **✨ Harvested**: Completed achievements that never disappear

### **ADHD-Friendly Features**

- **Executive Function Support**: External scaffolding for planning, focus, and
  task management
- **Knowledge Persistence**: Never lose learning when switching between
  interests
- **AI-Powered Breakdown**: Transform overwhelming projects into manageable
  25-minute quests
- **Inventory Management**: Track supplies and materials to prevent duplicate
  purchases
- **Body Doubling**: Virtual co-working for focus and accountability
- **Complete Data Sovereignty**: Self-hosted with full control over your
  personal data

---

## 🚀 Quick Start

### Prerequisites

- Docker and Docker Compose installed
- 8GB+ RAM available
- 50GB+ storage space (Less if not using local AI features)
- Optional: GPU for enhanced AI features

### Installation

1. **Clone the repository**

   ```bash
   git clone https://github.com/rghsoftware/nqi.git
   cd nqi
   ```

2. **Configure environment**

   ```bash
   cp .env.example .env
   # Edit .env with your preferred settings
   ```

3. **Deploy with Docker Compose**

   ```bash
   docker-compose up -d
   ```

4. **Access your NQI instance**
   - Web interface: `http://localhost:3000`
   - Mobile apps: Available for iOS and Android
   - Desktop apps: Available for Windows, macOS, and Linux

5. **Complete setup wizard**
   - Create your first Epic
   - Configure accessibility preferences
   - Set up optional AI features

---

## 🏗️ Architecture Overview

### **Technology Stack**

- **Frontend**: Flutter (cross-platform native performance)
- **Backend**: FastAPI (Python 3.11+)
- **Database**: PostgreSQL 15+ with pgvector extension
- **AI Processing**: Ollama (local LLM deployment)
- **Caching**: Redis
- **Storage**: S3-compatible object storage
- **Deployment**: Docker Compose

### **Key Design Principles**

- **Mobile-first**: Optimized for on-the-go task management
- **Offline-capable**: Local SQLite cache with conflict-free sync
- **Privacy-focused**: All AI processing happens locally
- **Accessibility-compliant**: WCAG 2.2 Level AA standards
- **Self-hosted**: Complete control over your data

---

## 🛠️ Development Setup

### **Local Development Environment**

1. **Backend Setup**

   ```bash
   cd backend
   python -m venv venv
   source venv/bin/activate  # On Windows: venv\Scripts\activate
   pip install -r requirements.txt
   python -m uvicorn main:app --reload
   ```

2. **Frontend Setup**

   ```bash
   cd frontend
   flutter pub get
   flutter run
   ```

3. **Database Setup**

   ```bash
   docker-compose -f docker-compose.dev.yml up postgres redis
   python manage.py migrate
   ```

### **Development Workflow**

- **Feature branches**: Create from `develop` branch
- **Testing**: Run full test suite before commits
- **Code quality**: Pre-commit hooks enforce formatting and linting
- **Documentation**: Update relevant docs with feature changes

---

## 📱 Platform Support

> [!NOTE]
> I currently do not own any Apple products so although macOS and iOS are
> supported by Flutter I currently have no way of testing or debugging those
> systems.
>
> Please still submit bug reports and we will get them solved.

### **Mobile Applications**

- **iOS 12+**: Native Flutter app with iOS-specific optimizations
- **Android 8+**: Material Design 3 compliance with Android adaptations

### **Desktop Applications**

- **Windows 10+**: Native desktop app with Windows-specific integrations
- **macOS 10.15+**: Native macOS app with system-level features
- **Linux**: Supports major distributions (Ubuntu 20.04+)

### **Web Application**

- **Progressive Web App**: Full offline capability
- **Browser Support**: Chrome 90+, Firefox 88+, Safari 14+, Edge 90+

---

## 🧠 Core Features

### **Quest-Based Agile Workflow**

Transform the traditional Agile methodology into a personal productivity system
that acknowledges the unique challenges of ADHD minds. Break overwhelming
projects into concrete, time-boxed quests that provide clear direction and
achievable progress.

### **Knowledge Management System**

Preserve learning and insights across interest cycles through structured
knowledge capture. Create searchable fragments of techniques, resources,
warnings, and context that persist beyond individual projects.

### **Inventory Tracking**

Visual, photo-based system for tracking project supplies and materials. Barcode
scanning, location tracking, and AI-powered duplicate detection prevent wasteful
repeat purchases.

### **Body Doubling Integration**

Three modes of virtual co-working support from silent presence to full video
sessions. Peer-to-peer connections maintain privacy while providing the
accountability benefits of working alongside others.

### **Accessibility & Customization**

Comprehensive accessibility features including reduced animations, high contrast
modes, large text options, and keyboard navigation. Fully customizable interface
accommodates different sensory needs and preferences.

---

## 🔧 Configuration Options

### **Accessibility Settings**

- **Reduce Animations**: Minimize motion for better focus
- **High Contrast**: Enhanced visual clarity
- **Large Text**: Scalable typography throughout the interface
- **Dark Mode**: Reduced eye strain for extended use

### **Privacy & Security**

- **Local AI Processing**: All suggestions generated on-device
- **Encrypted Knowledge Base**: Personal insights secured with end-to-end encryption
- **Configurable Sync**: Control what data synchronizes across devices
- **Anonymous Usage Analytics**: Optional and completely anonymized

### **Productivity Features**

- **Pomodoro Integration**: Built-in timer for focused work sessions
- **Gamification Elements**: XP, achievements, and streak tracking
- **Energy Level Tracking**: Pattern recognition for optimal task timing
- **Flexible Sprint Lengths**: Customize focus cycles to match your rhythms

---

## 🤝 Contributing

We welcome contributions from developers who understand the neurodivergent
experience and want to create better tools for our community.

### **Getting Started**

1. Read our Contributing Guidelines
2. Review the Code of Conduct
3. Check open issues tagged with `good-first-issue`
4. Submit pull requests with clear descriptions

### **Development Priorities**

- **Accessibility improvements**: WCAG compliance and neurodivergent-specific needs
- **Mobile optimizations**: Performance and battery efficiency
- **AI model improvements**: Better task breakdown and suggestion algorithms
- **Integration capabilities**: Calendar, task management, and communication tools

### **Bug Reports & Feature Requests**

Use GitHub issues to provide structured feedback. Include specific use cases and
accessibility considerations when relevant.

---

## 📚 Documentation

Documentation will be maintained in the `docs/` directory and will include user
guides, technical documentation, and deployment instructions as the project develops.

---

## 🔒 Privacy & Security

### **Data Sovereignty**

Your data remains under your complete control through self-hosting. No
third-party services process your personal information, project details, or
behavioral patterns.

### **Local AI Processing**

All artificial intelligence features run locally using Ollama, ensuring your
project details and personal insights never leave your infrastructure.

### **Encryption Standards**

- **Data at rest**: AES-256 encryption for database and file storage
- **Data in transit**: TLS 1.3 for all network communications
- **Personal insights**: End-to-end encryption for knowledge base entries

---

## 📊 System Requirements

### **Minimum Configuration**

- **CPU**: 4 cores
- **RAM**: 8GB
- **Storage**: 100GB SSD (This accounts for the OS installation as well)
- **Network**: Standard broadband connection

### **Recommended Configuration**

- **CPU**: 8 cores
- **RAM**: 16GB
- **GPU**: RTX 4060 Ti 16GB (for enhanced AI features)
- **Storage**: 500GB NVMe SSD
- **Network**: High-speed internet for initial setup and sync

### **Deployment Options**

- **Home Server**: Repurposed hardware or dedicated mini PC
- **VPS**: Cloud virtual private server (€3.79/month minimum)
- **Hybrid**: Local processing with cloud synchronization backup

---

## 🎯 Roadmap

### **Phase 1: Foundation** (12 weeks)

- Core Quest Board functionality
- Mobile application (iOS/Android)
- Basic AI integration
- Self-hosting infrastructure

### **Phase 2: Intelligence** (8 weeks)

- Advanced task breakdown algorithms
- Enhanced knowledge persistence
- Improved gamification systems
- Performance optimizations

### **Phase 3: Community** (8 weeks)

- Body doubling platform
- Template sharing marketplace
- Advanced collaboration features
- Plugin ecosystem

### **Phase 4: Ecosystem** (8 weeks)

- Calendar and communication integrations
- Wearable device support
- Advanced analytics and insights
- Enterprise deployment options

---

## 📄 License

Neural Quest Interface is licensed under the **GNU Affero General Public License v3.0**
(AGPL-3.0).

This copyleft license ensures that NQI remains free and open-source, with any
modifications or network-based services also required to be open-source. This
protects the neurodivergent community's access to these tools while encouraging
collaborative improvement.

See [LICENSE](LICENSE) for full license text.

---

## 🙏 Acknowledgments

This project builds upon extensive research in ADHD management, executive
function support, and neurodivergent-friendly design principles. We acknowledge
the scholars, clinicians, and advocates whose work makes tools like this possible.

---

## 📞 Support

For issues, feature requests, or questions, please use the GitHub issue tracker.
Documentation and setup guides will be maintained in the repository as the
project develops.
